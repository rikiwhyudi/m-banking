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

type customerHandlerImpl struct {
	customerUsecase ports.CustomerUsecase
	validation      *validator.Validate
	wg              sync.WaitGroup
}

func NewCustomerHandler(customerUsecase ports.CustomerUsecase) *customerHandlerImpl {
	return &customerHandlerImpl{customerUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *customerHandlerImpl) RegisterCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewDecoder(r.Body).Decode(&response)
		return
	}

	// Validate request input using go-playground/validator.
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		customerResponse, err := h.customerUsecase.RegisterCustomerUsecase(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: customerResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()

}
