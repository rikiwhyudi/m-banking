package customerdto

type CustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	Nik         int    `json:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
