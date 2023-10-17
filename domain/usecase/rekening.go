package usecase

import (
	"context"
	accNumberdto "m-banking/dto/rekening"
)

type AccountNumberUsecase interface {
	GetBalanceUsecase(ctx context.Context, accountNumber int) (*accNumberdto.AccountNumberResponse, error)
	DepositUsecase(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
	CashoutUsecase(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
	TransferUsecase(ctx context.Context, transfer accNumberdto.TransferBalanceRequest) (*accNumberdto.TransferBalanceResponse, error)
}
