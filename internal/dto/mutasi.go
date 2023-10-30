package dto

import "time"

type TransactionRequest struct {
	AccountNumberID int       `json:"id"`
	TransactionCode string    `json:"transaction_code"`
	Amount          float64   `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
}
