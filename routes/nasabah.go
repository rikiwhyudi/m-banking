package routes

import (
	"m-banking/internal/adapters/delivery/http"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/usecase"
	"m-banking/pkg/database/sql"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {

	customerRepository := repository.NewCustomerRepository(sql.DB)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	h := http.NewCustomerHandler(customerUsecase)

	r.HandleFunc("/daftar", h.RegisterCustomerHandler).Methods("POST")
}
