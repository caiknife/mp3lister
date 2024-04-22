package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
	"github.com/caiknife/mp3lister/lib/pay/rustore/payments"
)

const (
	packageName       = "com.my.app" // packageName вашего приложения, например com.my.app
	purchaseToken     = "312.123"    // purchaseToken вида "invoice_id.user_id"
	subscriptionToken = "123.321"    // subscriptionToken вида "invoice_id.user_id"
	subscriptionID    = "promo"      // код продукта-подписки. Указывается разработчиком при создании продукта в RuStore Консоли.
)

type Options struct {
	PrivateKey   string `json:"privateKey"`
	PrivateKeyID string `json:"privateKeyID"`
	CompanyID    string `json:"companyID"`
}

var file = flag.String(
	"file",
	"",
	"file with options: privateKey, privateKeyID, companyID",
)
var options Options

func main() {
	flag.Parse()

	// загружаем ключ, id ключа и companyID

	data, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(
		data,
		&options,
	); err != nil {
		panic(err)
	}

	// создали нового клиента для работы с платежами:

	cl := client.New(
		options.PrivateKeyID,
		options.PrivateKey,
		options.CompanyID,
	)
	if err = cl.Auth(); err != nil {
		panic(err)
	}

	pay := payments.New(cl, packageName)
	ctx, cancel := context.WithTimeout(context.Background(), client.TimeOutSeconds*10*time.Second)

	defer cancel()

	// получить информацию по платежу с помощью purchaseToken:

	paymentInfo, err := pay.GetProductionPaymentInfo(
		ctx,
		purchaseToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nИнформация по платежу:%+v\n", paymentInfo)

	// получить информацию по подписке с помощью subscriptionToken:

	subscriptionInfo, err := pay.GetSubscriptionInfo(
		ctx,
		subscriptionToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nИнформация по подписке:%+v\n", subscriptionInfo)

	// получить информацию по подписке V2:

	subscriptionInfoV2, err := pay.GetSubscriptionInfoV2(
		ctx,
		packageName,
		subscriptionID,
		subscriptionToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nИнформация по подписке V2:%+v\n", subscriptionInfoV2)

	// получить информацию по подписке V3:

	subscriptionInfoV3, err := pay.GetSubscriptionInfoV3(
		ctx,
		packageName,
		subscriptionID,
		subscriptionToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nИнформация по подписке V3:%+v\n", subscriptionInfoV3)

	// получить статус подписки:

	subscriptionState, err := pay.GetSubscriptionState(
		ctx,
		subscriptionToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nСтатус подписки:%+v\n", subscriptionState)

	// подтвердить получение подписки:

	submitSubscription, err := pay.SubmitSubscription(
		ctx,
		packageName,
		subscriptionID,
		subscriptionToken,
	)
	if submitSubscription.Body != nil {
		log.Fatal(err)
	} // при успешном ответе Body ответа пустой

	fmt.Printf("\nПодтвердили получение подписки:%+v\n", submitSubscription)

}
