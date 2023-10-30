package models

import "time"

type AccountNumber struct {
	ID            int           `json:"id" gorm:"auto_increment"`
	AccountNumber int           `json:"account_number" gorm:"type: bigint;unique"`
	CustomerID    int           `json:"-" gorm:"foreignKey:CustomerID"` // has many fields from customer
	Balance       float64       `json:"balance" gorm:"type: decimal(12,2)"`
	Transaction   []Transaction `json:"-" gorm:"foreignKey:AccountNumberID"` //has many
	CreatedAt     time.Time     `json:"created_at"`
}
