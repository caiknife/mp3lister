package pay

import (
	"testing"

	"github.com/caiknife/mp3lister/config"
)

func TestRuStoreAPI(t *testing.T) {
	store := NewRuStore("", config.Pay.RuStore.CompanyID, config.Pay.RuStore.PrivateKey, config.Pay.RuStore.PackageName)
	info, err := store.GetPurchaseInfo(config.Pay.RuStore.PurchaseToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
	t.Log(store.CheckStatus(info))
}
