package web

import (
	"github.com/go-resty/resty/v2"
)

func Get(url string, param map[string]string) *resty.Response {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(param).
		SetHeader("Accept", "application/json").
		Get(url)
	if err != nil {
		// todo 待處理
	}
	return resp
}

func Post(url string, body any, result any) *resty.Response {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&result). // or SetResult(AuthSuccess{}).
		Post(url)
	if err != nil {
		// todo 待處理
	}
	return resp
}
