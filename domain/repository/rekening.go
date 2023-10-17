package repository

import (
	"context"
	models "m-banking/domain/models"
)

type AccountNumberRepository interface {
	GetBalanceRepository(ctx context.Context, accountNumber int) (models.AccountNumber, error)
	DepositRepository(ctx context.Context, deposit models.AccountNumber) (models.AccountNumber, error)
	CashoutRepository(ctx context.Context, cashout models.AccountNumber) (models.AccountNumber, error)
	TransferRepository(ctx context.Context, sender models.AccountNumber, receiver models.AccountNumber) (models.AccountNumber, error)
}
