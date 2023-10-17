package routes

import (
	"m-banking/mutasi/delivery/http"
	"m-banking/mutasi/repository"
	"m-banking/mutasi/usecase"
	"m-banking/pkg/postgresql"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {

	transactionRepository := repository.NewTransactionRepositoryImpl(postgresql.DB)
	transactionUsecase := usecase.NewTransactionUsecaseImpl(transactionRepository)

	h := http.NewTransactionHandlerImpl(transactionUsecase)

	r.HandleFunc("/mutasi/{id}", h.GetTransactionHandler).Methods("GET")

}
