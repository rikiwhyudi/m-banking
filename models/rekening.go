package models

type AccountNumber struct {
	ID            int           `json:"account_number_id" gorm:"primary_key:auto_increment"`
	AccountNumber int           `json:"account_number" gorm:"type: bigint;unique"`
	CustomerID    int           `json:"-" gorm:"foreignKey:CustomerID"` // has many fields from customer
	Balance       float64       `json:"balance" gorm:"type: decimal(12,2)"`
	Transaction   []Transaction `json:"-" gorm:"foreignKey:AccountNumberID"` //has many
}
