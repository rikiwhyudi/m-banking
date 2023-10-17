package models

type Customer struct {
	ID            int             `json:"account_id" gorm:"primary_key:auto_increment"`
	Name          string          `json:"name"  gorm:"type: varchar(255)"`
	Nik           int             `json:"nik" gorm:"type: bigint;unique"`
	PhoneNumber   string          `json:"phone_number" gorm:"type: varchar(255);unique"`
	AccountNumber []AccountNumber `json:"bank_account" gorm:"foreignKey:CustomerID"` //has many
}
