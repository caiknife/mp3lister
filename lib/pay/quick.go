package pay

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/caiknife/mp3lister/lib/fjson"
)

const (
	checkUserInfoURL = "http://checkuser.quickapi.net/v2/checkUserInfo"
)

type QuickClient struct {
	MD5Key      string `json:"md5_key"`
	CallbackKey string `json:"callback_key"`
}

func NewQuickClient(md5Key, callbackKey string) *QuickClient {
	c := &QuickClient{
		MD5Key:      md5Key,
		CallbackKey: callbackKey,
	}
	return c
}

type QuickCheckUserInfo struct {
	Token       string `json:"token"`
	UID         string `json:"uid"`
	ProductCode string `json:"product_code"`
	ChannelCode string `json:"channel_code"`
}

func (c *QuickCheckUserInfo) String() string {
	toString, _ := fjson.MarshalToString(c)
	return toString
}

func (q *QuickClient) VerifyUser(c *QuickCheckUserInfo) string {
	v := url.Values{}
	v.Set("token", c.Token)
	v.Set("uid", c.UID)
	v.Set("product_code", c.ProductCode)
	v.Set("channel_code", c.ChannelCode)
	v.Encode()
	reqURL := checkUserInfoURL + "?" + v.Encode()
	fmt.Println(reqURL)
	get, err := resty.New().R().Get(reqURL)
	if err != nil {
		return ""
	}
	return get.String()
}

func (q *QuickClient) VerifyOrder(ndData, sign, md5Sign string) bool {
	d := fmt.Sprintf("%s%s%s", ndData, sign, q.MD5Key)
	return cryptor.Md5String(d) == md5Sign
}

func (q *QuickClient) Decode(ndData string) (i *QuickCheckUserInfo, err error) {
	data := decryptData(ndData, q.CallbackKey)
	i = &QuickCheckUserInfo{}
	err = xml.Unmarshal([]byte(data), i)
	if err != nil {
		err = errors.WithMessage(err, "xml unmarshal")
		return nil, err
	}
	return i, nil
}

func encryptData(codeData string, callbackKey string) string {
	dataArr := []rune(codeData)
	keyArr := []byte(callbackKey)
	keyLen := len(keyArr)

	var tmpList []int

	for index, value := range dataArr {
		base := int(value)
		dataString := base + int(0xFF&keyArr[index%keyLen])
		tmpList = append(tmpList, dataString)
	}

	var str string

	for _, value := range tmpList {
		str += "@" + fmt.Sprintf("%d", value)
	}
	return str
}

func decryptData(ntData string, callbackKey string) string {
	strLen := len(ntData)
	newData := []rune(ntData)
	resultData := string(newData[1:strLen])
	dataArr := strings.Split(resultData, "@")
	keyArr := []byte(callbackKey)
	keyLen := len(keyArr)

	var tmpList []int

	for index, value := range dataArr {
		base, _ := strconv.Atoi(value)
		dataString := base - int(0xFF&keyArr[index%keyLen])
		tmpList = append(tmpList, dataString)
	}

	var str string

	for _, val := range tmpList {
		str += string(rune(val))
	}
	return str
}

type Message struct {
	IsTest       string `json:"is_test" xml:"is_test"`
	Channel      string `json:"channel" xml:"channel"`
	ChannelUID   string `json:"channel_uid" xml:"channel_uid"`
	GameOrder    string `json:"game_order" xml:"game_order"`
	OrderNo      string `json:"order_no" xml:"order_no"`
	PayTime      string `json:"pay_time" xml:"pay_time"`
	Amount       string `json:"amount" xml:"amount"`
	Status       string `json:"status" xml:"status"`
	ExtrasParams string `json:"extras_params" xml:"extras_params"`
}

type QuicksdkMessage struct {
	XMLName xml.Name `xml:"quicksdk_message"`
	Message Message  `xml:"message"`
}

func (n *QuicksdkMessage) String() string {
	toString, _ := fjson.MarshalToString(n)
	return toString
}

func (n *QuicksdkMessage) XML() string {
	marshal, err := xml.Marshal(n)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (n *QuicksdkMessage) IsTest() bool {
	return n.Message.IsTest == "1"
}

func (n *QuicksdkMessage) Success() bool {
	return n.Message.Status == "0"
}

func (n *QuicksdkMessage) GameOrder() string {
	return n.Message.GameOrder
}

func (n *QuicksdkMessage) OrderNo() string {
	return n.Message.OrderNo
}
