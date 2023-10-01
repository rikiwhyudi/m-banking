package routes

import (
	"e-wallet/handlers"
	"e-wallet/pkg/postgresql"
	"e-wallet/repositories"
	"e-wallet/service"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {

	transactionRepository := repositories.NewRepositoryTransactionImpl(postgresql.DB)
	transactionService := service.NewServiceTransactionImpl(transactionRepository)

	h := handlers.NewHandlerTransactionImpl(transactionService)

	r.HandleFunc("/mutasi/{id}", h.GetTransactionHandler).Methods("GET")

}
