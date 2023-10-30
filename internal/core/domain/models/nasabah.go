package models

import "time"

type Customer struct {
	ID            int             `json:"id" gorm:"primary_key:auto_increment"`
	Name          string          `json:"name"  gorm:"type: varchar(255)"`
	Nik           int             `json:"nik" gorm:"type: bigint;unique"`
	PhoneNumber   string          `json:"phone_number" gorm:"type: varchar(255);unique"`
	UserID        int             `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User          User            `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccountNumber []AccountNumber `json:"bank_account" gorm:"foreignKey:CustomerID"` //has many
	CreatedAt     time.Time       `json:"created_at"`
}
