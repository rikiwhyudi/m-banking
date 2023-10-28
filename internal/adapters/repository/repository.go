package repository

import "gorm.io/gorm"

type repositoriesImpl struct {
	db *gorm.DB
}
