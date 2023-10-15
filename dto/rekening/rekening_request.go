package accNumberdto

type AccountNumberRequest struct {
	AccountNumber int     `json:"account_number" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

type TransferBalanceRequest struct {
	SenderAccountNumber   int     `json:"sender_account_number" validate:"required"`
	ReceiverAccountNumber int     `json:"receiver_account_number" validate:"required"`
	Amount                float64 `json:"amount" validate:"required"`
}
