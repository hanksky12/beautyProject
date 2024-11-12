package web

import (
	"github.com/go-resty/resty/v2"
)

func Get(url string, param map[string]string) {
	client := resty.New()
	// todo 待處理

	resp, err := client.R().
		SetQueryParams(param).
		SetHeader("Accept", "application/json").
		Get(url)
}

func Post(url string, body any, result any) {
	client := resty.New()
	// todo 待處理
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&result). // or SetResult(AuthSuccess{}).
		Post(url)
}
