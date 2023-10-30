package dto

type CustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	Nik         int    `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
}

type CustomerResponse struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Nik           int         `json:"nik"`
	Email         string      `json:"email"`
	PhoneNumber   string      `json:"phone_number"`
	Status        string      `json:"status"`
	AccountNumber interface{} `json:"bank_account"`
}
