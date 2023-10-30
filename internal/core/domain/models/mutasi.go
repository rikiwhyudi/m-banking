package models

import "time"

type Transaction struct {
	ID              int       `json:"id" gorm:"primary_key:auto_increment"`
	AccountNumberID int       `json:"-" gorm:"foreignKey:AccountNumberID"` // has many fields from account_number
	TransactionCode string    `json:"transaction_code" gorm:"type: varchar(255)"`
	Amount          float64   `json:"amount" gorm:"type: decimal(12,2)"`
	CreatedAt       time.Time `json:"created_at"`
}
