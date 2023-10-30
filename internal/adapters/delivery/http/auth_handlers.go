package http

import (
	"context"
	"encoding/json"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type authHandlerImpl struct {
	authUsecase ports.AuthUsecase
	validation  *validator.Validate
	wg          sync.WaitGroup
}

func NewAuthHandler(authUsecase ports.AuthUsecase) *authHandlerImpl {
	return &authHandlerImpl{authUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *authHandlerImpl) RegisterAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		authRegisterResponse, err := h.authUsecase.RegisterAuthUsecase(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: authRegisterResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}

func (h *authHandlerImpl) LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		authLoginResponse, err := h.authUsecase.LoginAuthUsecase(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: authLoginResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}
