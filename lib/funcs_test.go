package lib

import (
	"testing"
)

const (
	m3uFile = "/Users/caiknife/20240424.m3u"
)

func TestGetOutputFile(t *testing.T) {
	t.Log(GetOutputFile(m3uFile, "csv"))
}

func TestWriteCSV(t *testing.T) {
	u, err := ReadM3U(m3uFile)
	if err != nil {
		t.Error(err)
		return
	}
	err = WriteCSV(u, GetOutputFile(m3uFile, "csv"))
	if err != nil {
		t.Error(err)
		return
	}
}

func TestReadM3U_FileNameSlice(t *testing.T) {
	u, err := readM3U(m3uFile)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)
	t.Log(u.Len())
}

func TestReadM3U_MP3Slice(t *testing.T) {
	u, err := ReadM3U(m3uFile)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(u)
	t.Log(u.Len())
}
