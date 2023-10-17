package routes

import (
	"m-banking/internal/rekening/delivery/http"
	"m-banking/internal/rekening/repository"
	"m-banking/internal/rekening/usecase"
	"m-banking/pkg/postgresql"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	accountNumberRepository := repository.NewAccountNumberRepositoryImpl(postgresql.DB)
	accountNumberUsecase := usecase.NewAccountNumberUsecaseImpl(accountNumberRepository)

	h := http.NewAccountNumberHandlerImpl(accountNumberUsecase)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
	r.HandleFunc("/transfer", h.TransferHandler).Methods("PATCH")
}
