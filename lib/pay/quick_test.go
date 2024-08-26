package pay

import (
	"encoding/xml"
	"testing"
)

func TestQuickClient_VerifyUser(t *testing.T) {
	c := NewQuickClient("", "")
	v := &QuickCheckUserInfo{
		Token:       "@178@83@173@158@157@88@108@86@118@98@117@107@105@106@108@99@104@120@108@103@112@123@125@106@96@101@104@110@104@115@105@101@169@187@175@156@163@183@152@164@101@155@134@217@160@158@157@87@115@103@102@105@101@99@105@99@99@110@105@92@85@154@157@152@165@163@158@163@121@151@90@112@90@100@102@87@157@151@219@196@217@215@134@165@121@163@225",
		UID:         "D2A864635A709FD302080B508FF98D49",
		ProductCode: "64345624204336603757759703868145",
		ChannelCode: "",
	}
	user := c.VerifyUser(v)
	t.Log(user)
}

const (
	inputXML = `<quicksdk_message>
<message>
<is_test>0</is_test>
<channel>8888</channel>
<channel_uid>231845</channel_uid>
<game_order>123456789</game_order>
<order_no>12520160612114220441168433</order_no>
<pay_time>2016-06-12 11:42:20</pay_time>
<amount>1.00</amount>
<status>0</status>
<extras_params>{1}_{2}</extras_params>
</message>
</quicksdk_message>`
)

func TestNtData_XML(t *testing.T) {
	d := &QuicksdkMessage{}
	t.Log(d.XML())

	err := xml.Unmarshal([]byte(inputXML), d)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(d.XML())
	t.Log(d)
}
