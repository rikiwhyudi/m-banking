package dto

type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"message"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
