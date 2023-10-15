package accNumberdto

type AccountNumberResponse struct {
	ID            int     `json:"account_number_id"`
	AccountNumber int     `json:"account_number"`
	Balance       float64 `json:"balance"`
}

type TransferBalanceResponse struct {
	Name           string  `json:"name"`
	SenderNumber   int     `json:"sender_account_number"`
	Balance        float64 `json:"balance"`
	ReceiverNumber int     `json:"receiver_account_number"`
	Amount         float64 `json:"amount"`
}
