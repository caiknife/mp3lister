package main

import (
	"github.com/caiknife/mp3lister/cmd/fcm/fcm"
	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/types"
)

var (
	AuthSuccess types.Slice[*fcm.Check]
	AuthNotYet  types.Slice[*fcm.Check]
	AuthFail    types.Slice[*fcm.Check]
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
	lib.InitJSONConfig(&AuthSuccess, "auth_success.json")
	lib.InitJSONConfig(&AuthNotYet, "auth_notyet.json")
	lib.InitJSONConfig(&AuthFail, "auth_fail.json")
	lib.InitJSONConfig(&QuerySuccess, "query_success.json")
	lib.InitJSONConfig(&QueryNotYet, "query_notyet.json")
	lib.InitJSONConfig(&QueryFail, "query_fail.json")
	lib.InitJSONConfig(&Report, "report.json")
}
