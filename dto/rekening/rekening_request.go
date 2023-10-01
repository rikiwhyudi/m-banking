package accNumberdto

type AccountNumberRequest struct {
	AccountNumber int     `json:"account_number" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}
