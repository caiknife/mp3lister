package publishing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	neturl "net/url"
	"os"
	"time"

	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
)

const (
	versionURL               = client.BaseURL + "/%s/version"
	versionDeleteURL         = client.BaseURL + "/%s/version/%d"
	apkFileUploadURL         = client.BaseURL + "/%s/version/%d/apk?isMainApk=%s"
	aabFileUploadURL         = client.BaseURL + "/%s/version/%d/aab"
	iconFileUploadURL        = client.BaseURL + "/%s/version/%d/image/icon"
	screenFileUploadURL      = client.BaseURL + "/%s/version/%d/image/screenshot/%s/%s"
	changePublishSettingsURL = client.BaseURL + "/%s/version/%d/publish-settings"
	sendToModerationURL      = client.BaseURL + "/%s/version/%d/commit"
	manualPublishURL         = client.BaseURL + "/%s/version/%d/publish"
)

type Version struct {
	client      *client.Client
	packageName string
}

// nolint: tagliatelle
type CreateDraftText struct {
	AppName          string `json:"appName,omitempty"`
	AppType          string `json:"appType,omitempty"`
	Categories       string `json:"categories,omitempty"`
	AgeLegal         string `json:"ageLegal,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty"`
	FullDescription  string `json:"fullDescription,omitempty"`
	WhatsNew         string `json:"whatsNew,omitempty"`
	ModerInfo        string `json:"moderInfo,omitempty"`
	PriceValue       string `json:"priceValue,omitempty"`
	PublishType      string `json:"publishType,omitempty"`
	PublishDateTime  string `json:"publishDateTime,omitempty"`
	PartialValue     string `json:"partialValue,omitempty"`
}

func New(c *client.Client, packageName string) *Version {
	return &Version{
		client:      c,
		packageName: packageName,
	}
}

// nolint: tagliatelle
type CreateVersionDraftResponse struct {
	Code      string    `json:"code"`
	Message   *string   `json:"message"`
	Body      int       `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *Version) CreateVersionDraft(ctx context.Context, draftText CreateDraftText) (CreateVersionDraftResponse, error) {
	draftTextJSON, err := json.Marshal(draftText)
	if err != nil {
		return CreateVersionDraftResponse{}, err
	}

	url := fmt.Sprintf(
		versionURL,
		c.packageName,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(draftTextJSON),
	)
	if err != nil {
		return CreateVersionDraftResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return CreateVersionDraftResponse{}, err
	}

	defer response.Body.Close()

	createVersionDraft := CreateVersionDraftResponse{}
	err = json.NewDecoder(response.Body).Decode(&createVersionDraft)

	return createVersionDraft, err
}

type DeleteVersionDraftResponse struct {
	Code      string    `json:"code"`
	Message   *string   `json:"message"` // Message приходит только в случае ошибки
	Timestamp time.Time `json:"timestamp"`
}

func (c *Version) DeleteVersionDraft(ctx context.Context, versionID int) (DeleteVersionDraftResponse, error) {
	url := fmt.Sprintf(
		versionDeleteURL,
		c.packageName,
		versionID,
	)
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		return DeleteVersionDraftResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return DeleteVersionDraftResponse{}, err
	}

	defer response.Body.Close()

	deleteVersionDraft := DeleteVersionDraftResponse{}
	err = json.NewDecoder(response.Body).Decode(&deleteVersionDraft)

	return deleteVersionDraft, err

}

// nolint: tagliatelle
type GetVersion struct {
	Code      string         `json:"code"`
	Message   *string        `json:"message"`
	Body      GetVersionBody `json:"body"`
	Timestamp time.Time      `json:"timestamp"`
}

// nolint: tagliatelle
type GetVersionBody struct {
	Content       []Content `json:"content"`
	PageNumber    int       `json:"pageNumber"`
	PageSize      int       `json:"pageSize"`
	TotalElements int       `json:"totalElements"`
	TotalPages    int       `json:"totalPages"`
}

// nolint: tagliatelle
type Content struct {
	VersionID        int        `json:"versionId"`
	AppName          string     `json:"appName"`
	AppType          string     `json:"appType"`
	VersionName      *string    `json:"versionName"`
	VersionCode      int        `json:"versionCode"`
	VersionStatus    string     `json:"versionStatus"`
	PublishType      string     `json:"publishType"`
	PublishDateTime  *time.Time `json:"publishDateTime"`
	SendDateForModer *time.Time `json:"sendDateForModer"`
	PartialValue     int        `json:"partialValue"`
	WhatsNew         *string    `json:"whatsNew"`
	PriceValue       int        `json:"priceValue"`
	Paid             bool       `json:"paid"`
}

func (c *Version) GetVersion(ctx context.Context, params neturl.Values) (GetVersion, error) {
	url := fmt.Sprintf(
		versionURL,
		c.packageName,
	)

	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return GetVersion{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetVersion{}, err
	}

	defer response.Body.Close()

	getVersion := GetVersion{}
	err = json.NewDecoder(response.Body).Decode(&getVersion)

	return getVersion, err
}

type UploadFile struct {
	Code      string    `json:"code"`
	Message   *string   `json:"message"` // Message приходит только в случае ошибки
	Timestamp time.Time `json:"timestamp"`
}

func (c *Version) UploadAPKFile(ctx context.Context, versionID int, fileName string,
	isMainAPK bool, params neturl.Values) (UploadFile, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return UploadFile{}, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return UploadFile{}, err
	}

	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return UploadFile{}, err
	}

	contentType := bodyWriter.FormDataContentType()

	bodyWriter.Close()

	url := fmt.Sprintf(
		apkFileUploadURL,
		c.packageName,
		versionID,
		isMainAPK,
	)

	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bodyBuf,
	)
	if err != nil {
		return UploadFile{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{contentType})
	if err != nil {
		return UploadFile{}, err
	}

	defer response.Body.Close()

	uploadAPKFile := UploadFile{}
	err = json.NewDecoder(response.Body).Decode(&uploadAPKFile)

	return uploadAPKFile, err
}

func (c *Version) UploadAABFile(ctx context.Context, versionID int, fileName string) (UploadFile, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return UploadFile{}, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return UploadFile{}, err
	}

	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return UploadFile{}, err
	}

	contentType := bodyWriter.FormDataContentType()

	bodyWriter.Close()

	url := fmt.Sprintf(
		aabFileUploadURL,
		c.packageName,
		versionID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bodyBuf,
	)
	if err != nil {
		return UploadFile{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{contentType})
	if err != nil {
		return UploadFile{}, err
	}

	defer response.Body.Close()

	uploadAABFile := UploadFile{}
	err = json.NewDecoder(response.Body).Decode(&uploadAABFile)

	return uploadAABFile, err
}

func (c *Version) UploadIcon(ctx context.Context, versionID int, fileName string) (UploadFile, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return UploadFile{}, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return UploadFile{}, err
	}

	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return UploadFile{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := fmt.Sprintf(
		iconFileUploadURL,
		c.packageName,
		versionID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bodyBuf,
	)

	if err != nil {
		return UploadFile{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{contentType})
	if err != nil {
		return UploadFile{}, err
	}

	defer response.Body.Close()

	uploadIcon := UploadFile{}
	err = json.NewDecoder(response.Body).Decode(&uploadIcon)

	return uploadIcon, err
}

func (c *Version) UploadScreenshot(ctx context.Context, versionID int, fileName string,
	orientation string, ordinal string) (UploadFile, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		return UploadFile{}, err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		return UploadFile{}, err
	}

	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return UploadFile{}, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	url := fmt.Sprintf(
		screenFileUploadURL,
		c.packageName,
		versionID,
		orientation,
		ordinal,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bodyBuf,
	)

	if err != nil {
		return UploadFile{}, err
	}

	request.Header.Set("Content-Type", contentType)

	response, err := c.client.Do(request, client.RequestOpts{contentType})
	if err != nil {
		return UploadFile{}, err
	}

	defer response.Body.Close()

	uploadScreenshot := UploadFile{}
	err = json.NewDecoder(response.Body).Decode(&uploadScreenshot)

	return uploadScreenshot, err

}

type PublishResponse struct {
	Code      string    `json:"code"`
	Message   *string   `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type CreatePublishSettingsText struct {
	PublishType     string `json:"publishType,omitempty"`
	PublishDateTime string `json:"publishDateTime,omitempty"`
	PartialValue    string `json:"partialValue,omitempty"`
}

func (c *Version) ChangePublishSettings(ctx context.Context, versionID int,
	publishSettingsText CreatePublishSettingsText) (PublishResponse, error) {
	publishSettingsTextJSON, err := json.Marshal(publishSettingsText)
	if err != nil {
		return PublishResponse{}, err
	}

	url := fmt.Sprintf(
		changePublishSettingsURL,
		c.packageName,
		versionID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(publishSettingsTextJSON),
	)
	if err != nil {
		return PublishResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return PublishResponse{}, err
	}

	defer response.Body.Close()

	changePublishSettingsURL := PublishResponse{}
	err = json.NewDecoder(response.Body).Decode(&changePublishSettingsURL)

	return changePublishSettingsURL, err

}

func (c *Version) SendToModeration(ctx context.Context, versionID int, params neturl.Values) (PublishResponse, error) {
	url := fmt.Sprintf(
		sendToModerationURL,
		c.packageName,
		versionID,
	)

	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		return PublishResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return PublishResponse{}, err
	}

	defer response.Body.Close()

	sendToModeration := PublishResponse{}
	err = json.NewDecoder(response.Body).Decode(&sendToModeration)

	return sendToModeration, err
}

func (c *Version) ManualPublish(ctx context.Context, versionID int) (PublishResponse, error) {
	url := fmt.Sprintf(
		manualPublishURL,
		c.packageName,
		versionID,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		return PublishResponse{}, err
	}

	response, err := c.client.Do(request, client.RequestOpts{})
	if err != nil {
		return PublishResponse{}, err
	}

	defer response.Body.Close()

	manualPublish := PublishResponse{}
	err = json.NewDecoder(response.Body).Decode(&manualPublish)

	return manualPublish, err

}
