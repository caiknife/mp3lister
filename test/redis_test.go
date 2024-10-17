package test

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/samber/lo"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

func TestRedisHIncr(t *testing.T) {
	result, err := config.RedisDefault.HIncrBy(context.TODO(), "myhash", "111111", 1).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	n := time.Now()
	next := datetime.EndOfDay(n)
	b, err := config.RedisDefault.Expire(context.TODO(), "myhash", next.Sub(n)).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(b)
}

func Test_GenMultiKey(t *testing.T) {
	for i := range 1000 {
		_, err := config.RedisDefault.Set(context.TODO(), fmt.Sprintf("test:key:%d", i), i, 0).Result()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func Test_RedisScan(t *testing.T) {
	var start uint64 = 0
	r := types.Slice[string]{}
	for {
		result, u, err := config.RedisDefault.Scan(context.TODO(), start, "test:key:*", 100).Result()
		if err != nil {
			t.Error(err)
			return
		}
		if u == 0 {
			break
		}
		start = u
		r = append(r, result...)
	}
	r.Sort(func(i, j int) bool {
		return r[i] < r[j]
	})
	t.Log(r)
}

const jsonStr = `[
            {
                "id": "294",
                "tag": "XT0X7E1R",
                "legion_name": "拖",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "462",
                "tag": "TG8VJ9R9",
                "legion_name": "嘎嘎嘎嘎嘎嘎",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "哈哈哈哈哈哈",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "465",
                "tag": "DB19WRHV",
                "legion_name": "发生发的",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "467",
                "tag": "N7ZSFDZE",
                "legion_name": "股份第三个",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "461",
                "tag": "PM49MO8H",
                "legion_name": "565656",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "295",
                "tag": "IGCWBRQX",
                "legion_name": "hgf",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "474",
                "tag": "5A0RUW3Q",
                "legion_name": "fdd",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "476",
                "tag": "96QCRZIE",
                "legion_name": "bb",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "555",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "479",
                "tag": "SC0R5BOH",
                "legion_name": "方撒",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "489",
                "tag": "9VM13YHH",
                "legion_name": "mmmmmmmmmmmm",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "458",
                "tag": "EYEG3T18",
                "legion_name": "阿达撒",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "490",
                "tag": "KWRXSOPP",
                "legion_name": "j s j s j s j s",
                "legion_banner": {
                    "symbol": "symbol_01",
                    "pattern": "AAA_11"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 1,
                "trophy_score": 14000,
                "slogan": "mmnn",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "457",
                "tag": "HRTUCBLC",
                "legion_name": "乐趣",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "491",
                "tag": "XFP750UN",
                "legion_name": "ffffff",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "456",
                "tag": "JMGPJHSP",
                "legion_name": "明年",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "493",
                "tag": "VZGQ78W5",
                "legion_name": "<i>888</i>",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "<i>99999999999999</i>",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "455",
                "tag": "E5HVMW3Z",
                "legion_name": "Fff",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "B",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "494",
                "tag": "8I3LXT9P",
                "legion_name": "金龟换酒",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 1,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "449",
                "tag": "CG4JFCEP",
                "legion_name": "1234567",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 1,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "495",
                "tag": "S9XGMR7R",
                "legion_name": "Qaaa",
                "legion_banner": {
                    "symbol": "symbol_05",
                    "pattern": "AAA_18"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "jjdjdjcjcjxjxkcjcjfhdjsjsksksjdhdksjsnsksksjdjsjsjshdkdidkdksksnccjdjsjndjdjdjdhcjdjdjdndjdjdjchcjci",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "444",
                "tag": "T16I46IJ",
                "legion_name": "Tony Zhang",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "497",
                "tag": "753XQOWG",
                "legion_name": "Uuj",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "437",
                "tag": "YSSXIYQ0",
                "legion_name": "No.3",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "420",
                "tag": "CCIIBA0C",
                "legion_name": "肯",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13898,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "429",
                "tag": "TRTI17ZW",
                "legion_name": "发胜多负少",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "364",
                "tag": "BCRN8374",
                "legion_name": "Aaaa",
                "legion_banner": {
                    "symbol": "symbol_04",
                    "pattern": "AG"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 1,
                "trophy_score": 13994,
                "slogan": "Ndjd xinsii sii sinisinjsinsjsjsjsjsjsjsjdjdjdjdjjddjdjdjjdisisisisisisjdjdjdjdjdhdjdjdjdjdjdjdjdjxh",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "395",
                "tag": "KTUV4RP5",
                "legion_name": "TOt",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "307",
                "tag": "1KJICYDD",
                "legion_name": "111111",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13982,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "389",
                "tag": "7Y209PMI",
                "legion_name": "Fate",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "310",
                "tag": "4FZDXSJS",
                "legion_name": "Uuju",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13982,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "366",
                "tag": "URMK6MEY",
                "legion_name": "Tony华为",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "269",
                "tag": "VRB8NI6I",
                "legion_name": "777777777777777",
                "legion_banner": {
                    "symbol": "symbol_01",
                    "pattern": "AAA_01"
                },
                "open_lvl": 0,
                "limit_trophy": 4000,
                "limit_week": 1,
                "trophy_score": 13976,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "353",
                "tag": "3S5MK13U",
                "legion_name": "999999",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "447",
                "tag": "P171BTG8",
                "legion_name": "ooo",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13958,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "334",
                "tag": "JMQVOMAL",
                "legion_name": "一个个",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "382",
                "tag": "3SUTMEBC",
                "legion_name": "哇",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13940,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "326",
                "tag": "564QOP31",
                "legion_name": "坎坎坷坷看坎坎坷坷看看看",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "470",
                "tag": "BXW07MDF",
                "legion_name": "Iisjsjsjsjsjajs",
                "legion_banner": {
                    "symbol": "symbol_05",
                    "pattern": "AM"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13939,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "322",
                "tag": "60YFRSZS",
                "legion_name": "mopping",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "350",
                "tag": "9DR2W5TR",
                "legion_name": "6666",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13928,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "308",
                "tag": "ZXDOPVN5",
                "legion_name": "fdfdsfsd",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "484",
                "tag": "YZ77UPLS",
                "legion_name": "M m m m m m KKk",
                "legion_banner": {
                    "symbol": "symbol_04",
                    "pattern": "SF_10"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13928,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "306",
                "tag": "Y3UQ1ULR",
                "legion_name": "Tony",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "292",
                "tag": "SFHFZ1RX",
                "legion_name": "ue je j ejej s",
                "legion_banner": {
                    "symbol": "symbol_01",
                    "pattern": "ABH"
                },
                "open_lvl": 1,
                "limit_trophy": 13000,
                "limit_week": 1,
                "trophy_score": 13916,
                "slogan": "记得记得记得记得记得记得记得就记得记得就js s j s s k m d d k k d n d n d j d j d j d j j d k d kdm d mdn ndb d bdn d j d",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "305",
                "tag": "IVQR9JHC",
                "legion_name": "76876786876",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 1,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "445",
                "tag": "A37IVVDR",
                "legion_name": "TonyPC",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 13910,
                "slogan": "",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "471",
                "tag": "MU36MSTR",
                "legion_name": "模拟卷",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "323",
                "tag": "HYKOTBXV",
                "legion_name": "Isisjsnjsjsnshs",
                "legion_banner": {
                    "symbol": "symbol_05",
                    "pattern": "AI"
                },
                "open_lvl": 0,
                "limit_trophy": 5000,
                "limit_week": 2,
                "trophy_score": 13850,
                "slogan": "Fsgsssdn dhfhrhrhdhdhdhdhdjudjdjdjddjdjdndjjddhdbbdbdbdbdbdbdbdddjdjsjissjjsjshdhdhhddhdbdhjdeiiosos",
                "join_war": 1,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "473",
                "tag": "F7OZBI49",
                "legion_name": "kkzz",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 0,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "en",
                "members_count": 1
            },
            {
                "id": "502",
                "tag": "8WGYUF9X",
                "legion_name": "v重新做v重新",
                "legion_banner": {
                    "symbol": "symbol_02",
                    "pattern": "AAA_02"
                },
                "open_lvl": 0,
                "limit_trophy": 1000,
                "limit_week": 12,
                "trophy_score": 14000,
                "slogan": "33333",
                "join_war": 0,
                "tank_num_limit": 0,
                "language": "zh",
                "members_count": 1
            }
        ]`

type e struct {
	ID       string `json:"id"`
	Language string `json:"language"`
}

var sl = types.Slice[*e]{}

func init() {
	_ = fjson.UnmarshalFromString(jsonStr, &sl)
}

func Test_Sort_SliceStable(t *testing.T) {
	language := "zh"
	strings := lo.Map[*e, string](sl, func(item *e, index int) string {
		return item.ID
	})
	t.Log(strings)
	sort.SliceStable(sl, func(i, j int) bool {
		if sl[i].Language == language && sl[j].Language == language {
			return false
		}
		if sl[i].Language == language && sl[j].Language != language {
			return true
		}
		if sl[i].Language != language && sl[j].Language == language {
			return false
		}
		if sl[i].Language != language && sl[j].Language != language {
			return false
		}
		return false
	})

	strings = lo.Map[*e, string](sl, func(item *e, index int) string {
		return item.ID
	})
	t.Log(strings)
}

func Test_Slice_SortStable(t *testing.T) {
	language := "zh"
	strings := lo.Map[*e, string](sl, func(item *e, index int) string {
		return item.ID
	})
	t.Log(strings)
	slices.SortStableFunc(sl, func(a, b *e) int {
		if a.Language == language && b.Language == language {
			return 0
		}
		if a.Language == language && b.Language != language {
			return -1
		}
		if a.Language != language && b.Language == language {
			return 1
		}
		if a.Language != language && b.Language != language {
			return 0
		}
		return 0
	})
	strings = lo.Map[*e, string](sl, func(item *e, index int) string {
		return item.ID
	})
	t.Log(strings)
}

func Benchmark_Sort_SliceStable(b *testing.B) {
	language := "zh"
	for i := 0; i < b.N; i++ {
		sort.SliceStable(sl, func(i, j int) bool {
			if sl[i].Language == language && sl[j].Language == language {
				return false
			}
			if sl[i].Language == language && sl[j].Language != language {
				return true
			}
			if sl[i].Language != language && sl[j].Language == language {
				return false
			}
			if sl[i].Language != language && sl[j].Language != language {
				return false
			}
			return false
		})
	}
}

func Benchmark_Slice_SortStable(b *testing.B) {
	language := "zh"
	for i := 0; i < b.N; i++ {
		slices.SortStableFunc(sl, func(a, b *e) int {
			if a.Language == language && b.Language == language {
				return 0
			}
			if a.Language == language && b.Language != language {
				return -1
			}
			if a.Language != language && b.Language == language {
				return 1
			}
			if a.Language != language && b.Language != language {
				return 0
			}
			return 0
		})
	}
}
