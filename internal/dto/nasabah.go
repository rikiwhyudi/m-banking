package dto

type CustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	Nik         int    `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type CustomerResponse struct {
	ID            int         `json:"account_id"`
	Name          string      `json:"name"`
	Nik           int         `json:"nik"`
	PhoneNumber   string      `json:"phone_number"`
	AccountNumber interface{} `json:"bank_account"`
}
