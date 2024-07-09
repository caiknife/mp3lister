package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
	"github.com/caiknife/mp3lister/lib/pay/rustore/publishing"
)

var (
	urlValues          = url.Values{}
	versionIDInt       = 123 // версия вашего черновика
	versionIDString    = "123"
	fileNameAPK        = "app-release.apk"
	fileNameIcon       = "icon.png"
	isMainAPK          = false
	fileNameAAB        = "app-release.aab"
	fileNameScreenshot = "testScreenshot.jpeg"
	draftText          = publishing.CreateDraftText{
		AppName: "Название2",
		AppType: "MAIN",
	}
	changePublishSettingsText = publishing.CreatePublishSettingsText{
		PartialValue: "10",
	}

	packageName = "com.app.name" // packageName вашего приложения
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

	// создали нового клиента для работы с версиями:

	cl := client.New(
		options.PrivateKeyID,
		options.PrivateKey,
		options.CompanyID,
	)
	err = cl.Auth()
	if err != nil {
		log.Fatal(err)
	}

	pub := publishing.New(
		cl,
		packageName,
	)
	ctx, cancel := context.WithTimeout(
		context.TODO(),
		client.TimeOutSeconds*10*time.Second,
	)

	defer cancel()

	// загрузить новый черновик версии:

	createVersionDraft, err := pub.CreateVersionDraft(
		ctx,
		draftText,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nСоздали черновик версии: %+v\n", createVersionDraft)

	// удалить черновик версии:

	deleteVersionDraft, err := pub.DeleteVersionDraft(
		ctx,
		versionIDInt,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nУдалили черновик версии: %+v\n", deleteVersionDraft)

	// получить все статусы версии приложений:

	getVersionStatus, err := pub.GetVersion(ctx, urlValues)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nполучили все статусы версий: %+v\n", getVersionStatus)

	// получить статусы версий приложений по id:
	urlValues.Set("id", versionIDString)
	getVersionStatusID, err := pub.GetVersion(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nполучили статус по id версии: %+v\n", getVersionStatusID)

	// получить статусы версий по page и size:
	urlValues = url.Values{}
	urlValues.Set("page", "0")
	urlValues.Set("size", "2")
	getVersionStatusPageSize, err := pub.GetVersion(
		ctx,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nполучили статус по page и size: %+v\n", getVersionStatusPageSize)

	// загрузить иконку:

	uploadIcon, err := pub.UploadIcon(
		ctx,
		versionIDInt,
		fileNameIcon,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nзагрузили  иконку: %+v\n", uploadIcon)

	// загрузить скриншот:

	uploadScreenshot, err := pub.UploadScreenshot(
		ctx,
		versionIDInt,
		fileNameScreenshot,
		"PORTRAIT",
		"1",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nзагрузили  скриншот: %+v\n", uploadScreenshot)

	// отправить приложение на модерацию:
	urlValues = url.Values{}
	urlValues.Set("priorityUpdate", "0")
	sendToModeration, err := pub.SendToModeration(
		ctx,
		versionIDInt,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nотправили приложение на модерацию: %+v\n", sendToModeration)

	// опубликовать версию вручную:
	manualPublish, err := pub.ManualPublish(
		ctx,
		versionIDInt,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nопубликовали версию вручную: %+v\n", manualPublish)

	// изменить настройки публикации:

	changePublishSettings, err := pub.ChangePublishSettings(
		ctx,
		versionIDInt,
		changePublishSettingsText,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nизменили настройки публикации: %+v\n", changePublishSettings)

	// загрузить APK файл:

	uploadAPK, err := pub.UploadAPKFile(
		ctx,
		versionIDInt,
		fileNameAPK,
		isMainAPK,
		urlValues,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nзагрузили  APK файл: %+v\n", uploadAPK)
	if uploadAPK.Message != nil {
		fmt.Println(*uploadAPK.Message)
	}

	// загрузить AAB файл:
	uploadAAB, err := pub.UploadAABFile(
		ctx,
		versionIDInt,
		fileNameAAB,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nзагрузили  AAB файл: %+v\n", uploadAAB)
	if uploadAAB.Message != nil {
		fmt.Println(*uploadAAB.Message)
	}

}
