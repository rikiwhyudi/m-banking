package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique"`
	Password  string    `json:"-" gorm:"type:varchar(255)"`
	Role      string    `json:"status" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
}
