package dto

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
