package pay

import (
	"testing"
)

func TestRuStoreAPI(t *testing.T) {
	store := NewRuStore(Pay.RuStore.KeyID, Pay.RuStore.CompanyID, Pay.RuStore.PrivateKey, Pay.RuStore.PackageName)
	info, err := store.GetPurchaseInfo(Pay.RuStore.PurchaseToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
	t.Log(store.CheckStatus(info))
}
