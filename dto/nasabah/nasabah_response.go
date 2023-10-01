package customerdto

type CustomerResponse struct {
	ID            int         `json:"account_id"`
	Name          string      `json:"name"`
	Nik           int         `json:"nik"`
	PhoneNumber   string      `json:"phone_number"`
	AccountNumber interface{} `json:"bank_account"`
}
