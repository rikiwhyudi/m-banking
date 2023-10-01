package accNumberdto

type AccountNumberResponse struct {
	ID            int     `json:"account_number_id"`
	AccountNumber int     `json:"account_number"`
	Balance       float64 `json:"balance"`
}
