package routes

import (
	"m-banking/internal/adapters/delivery/http"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/usecase"
	"m-banking/pkg/database/sql"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {

	authRepository := repository.NewAuthRepository(sql.DB)
	authUsecase := usecase.NewAuthUsecase(authRepository)

	h := http.NewAuthHandler(authUsecase)

	r.HandleFunc("/register", h.RegisterAuthHandler).Methods("POST")
	r.HandleFunc("/login", h.LoginAuthHandler).Methods("GET")
}
