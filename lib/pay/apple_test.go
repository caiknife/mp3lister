package pay

import (
	"context"
	"testing"
)

func TestNewApple(t *testing.T) {
	apple, err := NewApple(Pay.Apple.IssID, Pay.Apple.BundleID, Pay.Apple.KeyID, "", Pay.Apple.IsProduction)
	if err != nil {
		t.Error(err)
		return
	}

	id, err := apple.LookUpOrderId(context.TODO(), "MT8X6VQHXX")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id)

	info, err := apple.GetTransactionInfo(context.TODO(), "2000000676385466")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
