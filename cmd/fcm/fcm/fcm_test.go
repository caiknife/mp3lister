package fcm

import (
	"testing"
)

var (
	mi = NewFangChenMi(SecretKey, AppID, BizID)
)

func TestNewFangChenMi_Sign(t *testing.T) {
	s := mi.Sign(`{"data":"fUhHjbnK2ZcHNImr1UScCs1OcVEoSDpXZhi06fD2ottI44EAwxH8ycEReKTGnxjVtgulCdV6kMNSJ9754fVeU/3M2Nhb/Dv3HkuRYrdGKBGCjPVCIGnR/SZI5fpTZgEml+t+DjOzFlxADsFBBb5QRYAg0DOOchbw/zl5uKTwwePLNS4Ow2hU4D3v3nwxxxan5ypIN6GmCIid84eyEkui5Za0CKXfUXMNNuXQmLm68jFp01bIhtylk9ydaSgkFQ=="}`)
	t.Log(s)
}
