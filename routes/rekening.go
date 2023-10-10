package routes

import (
	"e-wallet/handlers"
	"e-wallet/pkg/postgresql"
	"e-wallet/repositories"
	"e-wallet/service"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	accountNumberRepository := repositories.NewRepositoryAccountNumberImpl(postgresql.DB)
	accountNumberService := service.NewServiceAccountNumberImpl(accountNumberRepository)

	h := handlers.NewHandlerAccountNumberImpl(accountNumberService)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
}
