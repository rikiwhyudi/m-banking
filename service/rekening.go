package service

import (
	accNumberdto "e-wallet/dto/rekening"
)

type AccountNumberService interface {
	GetBalanceService(accountNumber int) (*accNumberdto.AccountNumberResponse, error)
	DepositService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
	CashoutService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error)
}
