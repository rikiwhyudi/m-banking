package repository

import (
	"context"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"

	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) ports.AuthRepository {
	return &repository{db}
}

func (r *repository) RegisterAuthRepository(ctx context.Context, user models.User) (models.User, error) {
	var err error

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return user, tx.Error
	}

	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, err
}

func (r *repository) LoginAuthRepository(ctx context.Context, email string) (models.User, error) {
	var err error
	var user models.User

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.First(&user, "email=?", email).Error
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, err
}
