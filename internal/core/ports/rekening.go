package ports

import (
	"context"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/dto"
)

type AccountNumberRepository interface {
	GetBalanceRepository(ctx context.Context, accountNumber int) (models.AccountNumber, error)
	DepositRepository(ctx context.Context, deposit models.AccountNumber) (models.AccountNumber, error)
	CashoutRepository(ctx context.Context, cashout models.AccountNumber) (models.AccountNumber, error)
	TransferRepository(ctx context.Context, sender models.AccountNumber, receiver models.AccountNumber) (models.AccountNumber, error)
}

type AccountNumberUsecase interface {
	GetBalanceUsecase(ctx context.Context, accountNumber int) (*dto.AccountNumberResponse, error)
	DepositUsecase(ctx context.Context, account dto.AccountNumberRequest) (*dto.AccountNumberResponse, error)
	CashoutUsecase(ctx context.Context, account dto.AccountNumberRequest) (*dto.AccountNumberResponse, error)
	TransferUsecase(ctx context.Context, transfer dto.TransferBalanceRequest) (*dto.TransferBalanceResponse, error)
}
