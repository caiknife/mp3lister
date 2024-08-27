package test

import (
	"testing"

	"github.com/levigross/grequests"
)

func TestGrequests_GET(t *testing.T) {
	get, err := grequests.Get("https://httpbin.org/get?hello=world", nil)
	if err != nil {
		t.Error(err)
		return
	}
	defer get.Close()
	t.Log(get.String())
}

func TestGrequests_POST(t *testing.T) {
	ro := &grequests.RequestOptions{
		Data: map[string]string{
			"hello": "world",
			"foo":   "bar",
		},
		JSON: map[string]string{
			"hello": "caiknife",
		},
	}
	post, err := grequests.Post("https://httpbin.org/post", ro)
	if err != nil {
		t.Error(err)
		return
	}
	defer post.Close()
	t.Log(post.String())
}
