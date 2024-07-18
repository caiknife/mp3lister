package main

import (
	"github.com/pkg/errors"
	"gorm.io/gen"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/wartankcn"
	"github.com/caiknife/mp3lister/orm/wartankcn/model"
)

func modifyDB() error {
	q := wartankcn.Q
	err := q.Transaction(func(tx *wartankcn.Query) error {
		// 处理 game center 数据
		num, err := tbGameCenter(tx)
		if err != nil {
			err = errors.WithMessage(err, "modify game center")
			return err
		}
		logger.ConsoleLogger.Info("处理", tx.WtGamecenter.TableName(), "影响行数", num)

		// 处理 legion 数据
		num, err = tbLegion(tx)
		if err != nil {
			err = errors.WithMessage(err, "modify legion")
			return err
		}
		logger.ConsoleLogger.Info("处理", tx.WtLegion.TableName(), "影响行数", num)

		// 处理 order 数据
		num, err = tbOrder(tx)
		if err != nil {
			err = errors.WithMessage(err, "modify order")
			return err
		}
		logger.ConsoleLogger.Info("处理", tx.WtOrder.TableName(), "影响行数", num)

		// 处理 player 数据
		num, err = tbPlayer(tx)
		if err != nil {
			err = errors.WithMessage(err, "modify player")
			return err
		}
		logger.ConsoleLogger.Info("处理", tx.WtPlayer.TableName(), "影响行数", num)
		return nil
	})
	if err != nil {
		err = errors.WithMessage(err, "modify db transaction")
		return err
	}
	return nil
}

func tbPlayer(tx *wartankcn.Query) (n int64, err error) {
	tbP := tx.WtPlayer
	result1, err := tbP.Where(tbP.LegionID.Neq(-1)).UpdateSimple(
		tbP.ID.Add(playerIDIncrement),
		tbP.LegionID.Add(legionIDIncrement),
	)
	if err != nil {
		err = errors.WithMessage(err, "update player")
		return 0, err
	}
	result2, err := tbP.Where(tbP.LegionID.Eq(-1)).UpdateSimple(
		tbP.ID.Add(playerIDIncrement),
	)
	if err != nil {
		err = errors.WithMessage(err, "update player")
		return 0, err
	}
	return result1.RowsAffected + result2.RowsAffected, nil
}

func tbOrder(tx *wartankcn.Query) (n int64, err error) {
	tbO := tx.WtOrder
	result, err := tbO.Where(tbO.PlayerID).UpdateSimple(
		tbO.PlayerID.Add(playerIDIncrement),
	)
	if err != nil {
		err = errors.WithMessage(err, "update order")
		return 0, err
	}
	return result.RowsAffected, nil
}

func columnMembers(legion *model.WtLegion) error {
	var members types.Slice[uint64]
	err := fjson.Unmarshal(legion.Members, &members)
	if err != nil {
		err = errors.WithMessage(err, "unmarshal members")
		return err
	}
	for i, member := range members {
		members[i] = member + playerIDIncrement
	}
	legion.Members, err = fjson.Marshal(members)
	if err != nil {
		err = errors.WithMessage(err, "marshal members")
		return err
	}
	return nil
}

func tbLegion(tx *wartankcn.Query) (n int64, err error) {
	tbL := tx.WtLegion
	// 处理members字段里的player id
	batch, err := tbL.Where(tbL.ID).FindInBatch(defaultPageSize, func(tx gen.Dao, batch int) error {
		return nil
	})
	if err != nil {
		err = errors.WithMessage(err, "find in batch")
		return 0, err
	}
	for _, legion := range batch {
		err = columnMembers(legion)
		if err != nil {
			err = errors.WithMessage(err, "column members")
			return 0, err
		}
	}
	err = tbL.Where(tbL.ID).Save(batch...)
	if err != nil {
		err = errors.WithMessage(err, "save legion")
		return 0, err
	}
	// 处理主键ID
	result, err := tbL.Where(tbL.ID).UpdateSimple(
		tbL.ID.Add(legionIDIncrement),
	)
	if err != nil {
		err = errors.WithMessage(err, "update legion")
		return 0, err
	}
	return result.RowsAffected, nil
}

func tbGameCenter(tx *wartankcn.Query) (n int64, err error) {
	tbGc := tx.WtGamecenter
	result, err := tbGc.Where(tbGc.PlayerID).UpdateSimple(
		tbGc.PlayerID.Add(playerIDIncrement),
	)
	if err != nil {
		err = errors.WithMessage(err, "update game center")
		return 0, err
	}
	return result.RowsAffected, nil
}
