package repositories

import (
	"context"
	"e-wallet/models"
)

type AccountNumberRepository interface {
	GetBalanceRepository(ctx context.Context, ccountNumber int) (models.AccountNumber, error)
	DepositRepository(ctx context.Context, deposit models.AccountNumber) (models.AccountNumber, error)
	CashoutRepository(ctx context.Context, cashout models.AccountNumber) (models.AccountNumber, error)
}
