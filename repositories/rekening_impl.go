package repositories

import (
	"e-wallet/models"

	"gorm.io/gorm"
)

func NewRepositoryAccountNumberImpl(db *gorm.DB) AccountNumberRepository {
	return &repository{db}
}

func (r *repository) GetBalanceRepository(accountNumber int) (models.AccountNumber, error) {
	var err error
	var account models.AccountNumber

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return account, tx.Error
	}

	err = tx.First(&account, "account_number=?", accountNumber).Error
	if err != nil {
		tx.Rollback()
		return account, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return account, err
	}

	return account, err
}

func (r *repository) DepositRepository(deposit models.AccountNumber) (models.AccountNumber, error) {
	var err error

	// start db transaction
	tx := r.db.Begin()

	defer func() {
		if r := recover(); nil != r {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return deposit, tx.Error
	}

	err = tx.Save(&deposit).Error
	if err != nil {
		tx.Rollback()
		return deposit, err
	}

	// commit db transaction
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return deposit, err
	}

	return deposit, err
}

func (r *repository) CashoutRepository(cashout models.AccountNumber) (models.AccountNumber, error) {
	var err error

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return cashout, tx.Error
	}

	err = tx.Save(&cashout).Error
	if err != nil {
		tx.Rollback()
		return cashout, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return cashout, err
	}

	return cashout, err
}
