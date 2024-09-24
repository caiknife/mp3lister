package pay

import (
	"fmt"
	"testing"
)

func TestRuStoreAPI(t *testing.T) {
	Pay.RuStore.ForEach(func(conf *rustoreConf, i int) {
		store := NewRuStore(conf.KeyID, conf.CompanyID, conf.PrivateKey, conf.PackageName)

		t.Run(fmt.Sprintf("%s production purchase", conf.KeyID), func(t *testing.T) {
			info, err := store.GetPurchaseInfo(conf.PurchaseToken)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(info)
			t.Log(store.CheckStatus(info))
		})

		t.Run(fmt.Sprintf("%s sandbox purchase", conf.KeyID), func(t *testing.T) {
			info, err := store.GetPurchaseInfo(conf.SandboxPurchaseToken)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(info)
			t.Log(store.CheckStatus(info))
		})
	})
}
