package test

import (
	"net/url"
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/caiknife/mp3lister/lib/types"
)

func TestResty_Get(t *testing.T) {
	client := resty.New()
	req := client.R()
	req.SetHeaders(types.Map[string]{
		"name":  "caiknife",
		"email": "caiknife@hotmail.com",
	})
	req.SetQueryParams(types.Map[string]{
		"name":  "caiknife",
		"email": "caiknife@hotmail.com",
	})
	resp, err := req.EnableTrace().Get("https://httpbin.org/get")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Status(), resp.StatusCode())
	t.Log(resp.String())
}

type formData struct {
	url.Values
	Name  string `json:"name" form:"name" url:"name"`
	Email string `json:"email" form:"email" url:"email"`
}

func TestResty_Post(t *testing.T) {
	client := resty.New()
	req := client.R()
	data := url.Values{}
	data.Set("name", "caiknife")
	data.Set("email", "caiknife@homail.com")
	req.SetHeaders(types.Map[string]{
		"Content-Type": "application/x-www-form-urlencoded",
	}).SetBody(data.Encode())
	resp, err := req.Post("https://httpbin.org/post")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.String())
}
