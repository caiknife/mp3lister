package fcm

import (
	"context"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	SecretKey = "232c256864bc62d72e2b0a134be40fac"
	AppID     = "3cade069d435476a866490d7fd40efa6"
	BizID     = "1101999999"
)

const (
	AuthURL      = "https://wlc.nppa.gov.cn/test/authentication/check"
	AuthQueryURL = "https://wlc.nppa.gov.cn/test/authentication/query"
	ReportURL    = "https://wlc.nppa.gov.cn/test/collection/loginout"
)

type FangChenMi struct {
	SecretKey string `json:"secret_key"`
	AppID     string `json:"app_id"`
	BizID     string `json:"biz_id"`
}

var (
	DefaultFangChenMi = NewFangChenMi(SecretKey, AppID, BizID)
)

func NewFangChenMi(secretKey, appID, bizID string) *FangChenMi {
	f := &FangChenMi{
		SecretKey: secretKey,
		AppID:     appID,
		BizID:     bizID,
	}
	return f
}

func (f *FangChenMi) Encrypt(request string) (s string) {
	encryptRequest := AES128GCMWithBase64(request, f.SecretKey)
	s = fmt.Sprintf(`{"data":"%s"}`, encryptRequest)
	return s
}

func (f *FangChenMi) Sign(encryptBody string, ts int64) (s string) {
	s = fmt.Sprintf("%sappId%sbizId%stimestamps%d%s", f.SecretKey, f.AppID, f.BizID, ts, encryptBody)
	s = cryptor.Sha256(s)
	return s
}

func (f *FangChenMi) Auth(code string, req *Check) error {
	ts := time.Now().UnixMilli()
	s := f.Encrypt(req.String())
	sign := f.Sign(s, ts)
	url := fmt.Sprintf("%s/%s", AuthURL, code)
	logger.ConsoleLogger.Infoln("请求地址", url)

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	post, err := resty.New().R().SetContext(timeout).SetHeaders(types.Map[string]{
		"Content-Type": "application/json;charset=utf-8",
		"appId":        f.AppID,
		"bizId":        f.BizID,
		"timestamps":   cast.ToString(ts),
		"sign":         sign,
	}).SetBody(s).Post(url)
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s request failed", url))
		logger.ConsoleLogger.Errorln(err)
		return err
	}
	logger.ConsoleLogger.Infoln(post.String())
	return nil
}

func (f *FangChenMi) Query(code string, req *Query) error {
	ts := time.Now().UnixMilli()
	s := f.Encrypt("")
	sign := f.Sign(s, ts)
	url := fmt.Sprintf("%s/%s?ai=%s", AuthQueryURL, code, req.Ai)
	logger.ConsoleLogger.Infoln("请求地址", url)

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	get, err := resty.New().R().SetContext(timeout).SetHeaders(types.Map[string]{
		"appId":      f.AppID,
		"bizId":      f.BizID,
		"timestamps": cast.ToString(ts),
		"sign":       sign,
	}).Get(url)
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s request failed", url))
		logger.ConsoleLogger.Errorln(err)
		return err
	}
	logger.ConsoleLogger.Infoln(get.String())
	return nil
}

func (f *FangChenMi) LoginOrOut(code string, req *Collections) error {
	ts := time.Now().UnixMilli()
	s := f.Encrypt(req.String())
	sign := f.Sign(s, ts)
	url := fmt.Sprintf("%s/%s", ReportURL, code)
	logger.ConsoleLogger.Infoln("请求地址", url)

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	post, err := resty.New().R().SetContext(timeout).SetHeaders(types.Map[string]{
		"Content-Type": "application/json;charset=utf-8",
		"appId":        f.AppID,
		"bizId":        f.BizID,
		"timestamps":   cast.ToString(ts),
		"sign":         sign,
	}).SetBody(s).Post(url)
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s request failed", url))
		logger.ConsoleLogger.Errorln(err)
		return err
	}
	logger.ConsoleLogger.Infoln(post.String())
	return nil
}
