package dto

import "time"

type TransactionRequest struct {
	AccountNumberID int       `json:"account_number_id"`
	TransactionCode string    `json:"transaction_code"`
	Amount          float64   `json:"amount"`
	Date            time.Time `json:"date"`
}
