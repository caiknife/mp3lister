package test

import (
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/caiknife/mp3lister/lib/types"
)

const (
	qqMusicApiAlbumSongListURL = "https://i.y.qq.com/v8/fcg-bin/fcg_v8_album_info_cp.fcg?platform=h5page&albummid=ALBUMMID&g_tk=938407465&uin=0&format=json&inCharset=utf-8&outCharset=utf-8&notice=0&platform=h5&needNewCode=1&_=1459961045571"
	qqMusicApiSingerInfoURL    = "https://u.y.qq.com/cgi-bin/musicu.fcg?format=json&loginUin=0&hostUin=0inCharset=utf8&outCharset=utf-8&platform=yqq.json&needNewCode=0&data=%7B%22comm%22%3A%7B%22ct%22%3A24%2C%22cv%22%3A0%7D%2C%22singer%22%3A%7B%22method%22%3A%22get_singer_detail_info%22%2C%22param%22%3A%7B%22sort%22%3A5%2C%22singermid%22%3A%22SINGERMID%22%2C%22sin%22%3A0%2C%22num%22%3A50%7D%2C%22module%22%3A%22music.web_singer_info_svr%22%7D%7D"
	qqMusicApiMusicURL         = "https://u.y.qq.com/cgi-bin/musicu.fcg?format=json&data={%22req_0%22:{%22module%22:%22vkey.GetVkeyServer%22,%22method%22:%22CgiGetVkey%22,%22param%22:{%22filename%22:[%22PREFIXSONGMIDSONGMID.SUFFIX%22],%22guid%22:%2210000%22,%22songmid%22:[%22SONGMID%22],%22songtype%22:[0],%22uin%22:%220%22,%22loginflag%22:1,%22platform%22:%2220%22}},%22loginUin%22:%220%22,%22comm%22:{%22uin%22:%220%22,%22format%22:%22json%22,%22ct%22:24,%22cv%22:0}}"
)

var (
	testAlbumMIDs = types.Slice[string]{
		"002S7UBf14A62i",
	}
	testSingerMIDs = types.Slice[string]{
		"002a7qwy2WuP81",
	}
	testMusicMIDs = types.Slice[string]{
		"914383",
		"0016JyRa2YwUg3",
	}
)

func TestQQMusicAPI_MusicInfo(t *testing.T) {
	testMusicMIDs.ForEach(func(s string, i int) {
		url := strings.ReplaceAll(qqMusicApiMusicURL, "SONGMID", s)
		url = strings.ReplaceAll(url, "PREFIX", "M800")
		url = strings.ReplaceAll(url, "SUFFIX", "mp3")
		res, err := resty.New().R().Get(url)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(res.String())
	})
}

func TestQQMusicAPI_AlbumSongList(t *testing.T) {
	testAlbumMIDs.ForEach(func(s string, i int) {
		url := strings.ReplaceAll(qqMusicApiAlbumSongListURL, "ALBUMMID", s)
		res, err := resty.New().R().Get(url)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(res.String())
	})
}

func TestQQMusicAPI_SingerInfo(t *testing.T) {
	testSingerMIDs.ForEach(func(s string, i int) {
		url := strings.ReplaceAll(qqMusicApiSingerInfoURL, "SINGERMID", s)
		res, err := resty.New().R().Get(url)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(res.String())
	})
}
