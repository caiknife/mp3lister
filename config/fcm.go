package config

import (
	"github.com/caiknife/mp3lister/cmd/fcm/fcm"
	"github.com/caiknife/mp3lister/lib/types"
)

var (
	AuthSuccess types.Slice[*fcm.Check]
	AuthNotYet  types.Slice[*fcm.Check]
	AuthFailed  types.Slice[*fcm.Check]
)

var (
	QuerySuccess types.Slice[*fcm.Query]
	QueryNotYet  types.Slice[*fcm.Query]
	QueryFail    types.Slice[*fcm.Query]
)

var (
	Report types.Slice[*fcm.Behavior]
)

func init() {
	InitJSONConfig(&AuthSuccess, "auth_success.json")
	InitJSONConfig(&AuthNotYet, "auth_notyet.json")
	InitJSONConfig(&AuthFailed, "auth_fail.json")
	InitJSONConfig(&QuerySuccess, "query_success.json")
	InitJSONConfig(&QueryNotYet, "query_notyet.json")
	InitJSONConfig(&QueryFail, "query_fail.json")
	InitJSONConfig(&Report, "report.json")
}
