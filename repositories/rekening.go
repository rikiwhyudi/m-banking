package repositories

import "e-wallet/models"

type AccountNumberRepository interface {
	GetBalanceRepository(accountNumber int) (models.AccountNumber, error)
	DepositRepository(deposit models.AccountNumber) (models.AccountNumber, error)
	CashoutRepository(cashout models.AccountNumber) (models.AccountNumber, error)
}
