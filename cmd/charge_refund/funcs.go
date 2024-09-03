package main

import (
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gen"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/wartankcn"
	"github.com/caiknife/mp3lister/orm/wartankcn/model"
)

const (
	defaultBatchSize = 500
)

func doRefund() error {
	orders, err := loadDB()
	if err != nil {
		err = errors.WithMessage(err, "load db")
		return err
	}
	playerTotalCharges, err := transformOrder(orders)
	if err != nil {
		err = errors.WithMessage(err, "transform order")
		return err
	}

	refund, err := chargeRefund(playerTotalCharges)
	if err != nil {
		err = errors.WithMessage(err, "charge refund")
		return err
	}

	err = saveToDB(refund)
	if err != nil {
		err = errors.WithMessage(err, "save refund")
		return err
	}

	return nil
}

func saveToDB(refund types.Hash[string, *ChargeRefund]) error {
	// 先清空charge_refund表
	tbChargeRefund := wartankcn.ChargeRefund
	_, err := tbChargeRefund.Unscoped().Where(tbChargeRefund.ID).Delete()
	if err != nil {
		err = errors.WithMessage(err, "delete charge refund")
		return err
	}

	chargeRefunds := lo.MapToSlice[string, *ChargeRefund, *model.ChargeRefund](refund, func(key string, value *ChargeRefund) *model.ChargeRefund {
		return &model.ChargeRefund{
			GameCenterID: value.GameCenterID,
			PlayerID:     value.PlayerID,
			TotalCharge:  value.TotalCharge,
			Diamonds:     value.Diamonds,
			Acquired:     false,
		}
	})
	err = tbChargeRefund.CreateInBatches(chargeRefunds, defaultBatchSize)
	if err != nil {
		err = errors.WithMessage(err, "create charge refund")
		return err
	}
	return nil
}

func getPriceByProductID(productID string) float64 {
	productID = getProductID(productID)
	if ShopCfg.HasKey(productID) {
		return ShopCfg[productID].Price
	}
	return 0
}

func orderToPlayerOrder(order *model.WtOrder) *PlayerOrder {
	p := &PlayerOrder{
		PlayerID:  order.PlayerID,
		BundleID:  order.BundleID,
		ProductID: getProductID(order.ProductID),
		Price:     getPriceByProductID(order.ProductID),
	}
	return p
}

func chargeRefund(charges types.Hash[int64, float64]) (types.Hash[string, *ChargeRefund], error) {
	tableGameCenter := wartankcn.WtGamecenter
	batch, err := tableGameCenter.Where(
		tableGameCenter.PlayerID.In(charges.Keys()...),
	).FindInBatch(defaultBatchSize, func(tx gen.Dao, batch int) error {
		return nil
	})
	if err != nil {
		err = errors.WithMessage(err, "table game center find in batch")
		return nil, err
	}
	result := lo.SliceToMap[*model.WtGamecenter, string, *ChargeRefund](batch, func(item *model.WtGamecenter) (string, *ChargeRefund) {
		cr := &ChargeRefund{
			PlayerID:     item.PlayerID,
			GameCenterID: item.ID,
			TotalCharge:  charges[item.PlayerID],
			Diamonds:     0,
			Acquired:     false,
		}
		cr.CalcDiamonds()
		return item.ID, cr
	})
	return result, nil
}

func transformOrder(orders types.Slice[*model.WtOrder]) (types.Hash[int64, float64], error) {
	playerOrders := types.Hash[int64, types.Slice[*PlayerOrder]]{}
	orders.ForEach(func(order *model.WtOrder, i int) {
		if playerOrders.HasKey(order.PlayerID) {
			playerOrders[order.PlayerID] = append(playerOrders[order.PlayerID], orderToPlayerOrder(order))
		} else {
			playerOrders[order.PlayerID] = types.Slice[*PlayerOrder]{orderToPlayerOrder(order)}
		}
	})

	playerTotalCharges := types.Hash[int64, float64]{}
	playerOrders.ForEach(func(i int64, t types.Slice[*PlayerOrder]) {
		playerTotalCharges[i] = lo.SumBy[*PlayerOrder, float64](t, func(item *PlayerOrder) float64 {
			return item.Price
		})
	})
	return playerTotalCharges, nil
}

func loadDB() (types.Slice[*model.WtOrder], error) {
	tbOrder := wartankcn.WtOrder
	batch, err := tbOrder.Where(
		tbOrder.OrderStatus.Eq(8),
	).FindInBatch(defaultBatchSize, func(tx gen.Dao, batch int) error {
		return nil
	})
	if err != nil {
		err = errors.WithMessage(err, "table order find in batch")
		return nil, err
	}
	return batch, nil
}
