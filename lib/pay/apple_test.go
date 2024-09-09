package pay

import (
	"context"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	apple2 "github.com/go-pay/gopay/apple"

	"github.com/caiknife/mp3lister/lib/fjson"
)

func TestApple_Transaction(t *testing.T) {
	toString, err := fileutil.ReadFileToString(Pay.Apple.PrivateKey)
	if err != nil {
		t.Error(err)
		return
	}

	apple, err := NewApple(Pay.Apple.IssID, Pay.Apple.BundleID, Pay.Apple.KeyID, toString, Pay.Apple.IsProduction)
	if err != nil {
		t.Error(err)
		return
	}

	info, err := apple.GetTransactionInfo(context.TODO(), Pay.Apple.TransID)
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
}

func TestApp_Order(t *testing.T) {
	toString, err := fileutil.ReadFileToString(Pay.Apple.PrivateKey)
	if err != nil {
		t.Error(err)
		return
	}

	apple, err := NewApple(Pay.Apple.IssID, Pay.Apple.BundleID, Pay.Apple.KeyID, toString, Pay.Apple.IsProduction)
	if err != nil {
		t.Error(err)
		return
	}

	id, err := apple.LookUpOrderId(context.TODO(), Pay.Apple.OrderID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id.SignedTransactions)
	for _, transaction := range id.SignedTransactions {
		ti := &apple2.TransactionsItem{}
		err := apple2.ExtractClaims(string(transaction), ti)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(fjson.MarshalToString(ti))
	}
}
