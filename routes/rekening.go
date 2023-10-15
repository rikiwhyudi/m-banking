package routes

import (
	"m-banking/handlers"
	"m-banking/pkg/postgresql"
	"m-banking/repositories"
	"m-banking/service"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	accountNumberRepository := repositories.NewRepositoryAccountNumberImpl(postgresql.DB)
	accountNumberService := service.NewServiceAccountNumberImpl(accountNumberRepository)

	h := handlers.NewHandlerAccountNumberImpl(accountNumberService)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
	r.HandleFunc("/transfer", h.TransferHandler).Methods("PATCH")
}
