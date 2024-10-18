package dto

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Total   int    `json:"total"`
	Data    any    `json:"data"`
}
