package routes

import (
	"m-banking/internal/nasabah/delivery/http"
	"m-banking/internal/nasabah/repository"
	"m-banking/internal/nasabah/usecase"
	"m-banking/pkg/postgresql"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {

	customerRepository := repository.NewCustomerRepositoryImpl(postgresql.DB)
	customerUsecase := usecase.NewCustomerUsecaseImpl(customerRepository)

	h := http.NewCustomerHandlerImpl(customerUsecase)

	r.HandleFunc("/daftar", h.RegisterCustomerHandler).Methods("POST")
}
