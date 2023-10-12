package repositories

import (
	"context"
	"e-wallet/models"
)

type AccountNumberRepository interface {
	GetBalanceRepository(ctx context.Context, accountNumber int) (models.AccountNumber, error)
	DepositRepository(ctx context.Context, deposit models.AccountNumber) (models.AccountNumber, error)
	CashoutRepository(ctx context.Context, cashout models.AccountNumber) (models.AccountNumber, error)
}
