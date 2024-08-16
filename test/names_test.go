package test

import (
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/slice"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type P struct {
	Name string `json:"name"`
}

func TestNamesJsonToTxt(t *testing.T) {
	data := types.Slice[P]{}
	lib.InitJSONConfig(&data, "names_en.json")
	err := writeToTxt("names_en.txt", data)
	if err != nil {
		t.Error(err)
		return
	}
}

func writeToTxt(fileName string, data types.Slice[P]) error {
	str := slice.Map[P, string](data, func(index int, item P) string {
		return item.Name
	})
	err := fileutil.WriteStringToFile(fileName, strings.Join(str, "\n"), false)
	if err != nil {
		return err
	}
	return nil
}

func loadToSlice(fileName string, data *types.Slice[P]) error {
	toString, err := fileutil.ReadFileToString(fileName)
	if err != nil {
		return err
	}
	split := strings.Split(toString, "\n")
	*data = slice.Map[string, P](split, func(index int, item string) P {
		return P{Name: item}
	})
	return nil
}

func writeToJson(fileName string, data types.Slice[P]) error {
	toString, err := fjson.MarshalToString(data)
	if err != nil {
		return err
	}
	err = fileutil.WriteStringToFile(fileName, toString, false)
	if err != nil {
		return err
	}
	return nil
}

func TestNamesTxtToJson(t *testing.T) {
	data := types.Slice[P]{}
	err := loadToSlice("names_cn.txt", &data)
	if err != nil {
		t.Error(err)
		return
	}
	err = writeToJson("names_cn.json", data)
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_getApiServers(t *testing.T) {
	t.Log(getApiServer(false))
	t.Log(getApiServer(true))
	t.Log(apiServers)
}

func Test_getBackendServer(t *testing.T) {
	t.Log(getBackendServer(false))
	t.Log(getBackendServer(true))
	t.Log(backendServers)
}
