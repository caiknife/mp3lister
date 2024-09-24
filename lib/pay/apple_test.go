package pay

import (
	"context"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/go-pay/gopay/apple"

	"github.com/caiknife/mp3lister/lib/fjson"
)

func TestApple_Transaction(t *testing.T) {
	Pay.Apple.ForEach(func(conf *appleConf, i int) {
		toString, err := fileutil.ReadFileToString(conf.PrivateKey)
		if err != nil {
			t.Error(err)
			return
		}

		appleClient, err := NewApple(conf.IssID, conf.BundleID, conf.KeyID, toString, conf.IsProduction)
		if err != nil {
			t.Error(err)
			return
		}

		info, err := appleClient.GetTransactionInfo(context.TODO(), conf.TransID)
		if err != nil {
			t.Error(err)
			return
		}
		transaction, err := info.DecodeSignedTransaction()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(fjson.MarshalToString(transaction))
	})
}

func TestApp_Order(t *testing.T) {
	Pay.Apple.ForEach(func(conf *appleConf, i int) {
		toString, err := fileutil.ReadFileToString(conf.PrivateKey)
		if err != nil {
			t.Error(err)
			return
		}

		appleClient, err := NewApple(conf.IssID, conf.BundleID, conf.KeyID, toString, conf.IsProduction)
		if err != nil {
			t.Error(err)
			return
		}

		id, err := appleClient.LookUpOrderId(context.TODO(), conf.OrderID)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(id.SignedTransactions)
		for _, transaction := range id.SignedTransactions {
			ti := &apple.TransactionsItem{}
			err := apple.ExtractClaims(string(transaction), ti)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(fjson.MarshalToString(ti))
		}
	})

}
