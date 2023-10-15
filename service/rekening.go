package service

import (
	"context"
	accNumberdto "m-banking/dto/rekening"
)

type AccountNumberService interface {
	GetBalanceService(ctx context.Context, accountNumber int) (*accNumberdto.AccountNumberResponse, error)
	DepositService(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
	CashoutService(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
	TransferService(ctx context.Context, transfer accNumberdto.TransferBalanceRequest) (*accNumberdto.TransferBalanceResponse, error)
}
