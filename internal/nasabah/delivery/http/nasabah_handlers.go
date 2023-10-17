package http

import (
	"context"
	"encoding/json"
	customerdto "m-banking/dto/nasabah"
	dto "m-banking/dto/result"
	"m-banking/interfaces/usecase"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type customerHandler struct {
	customerUsecase usecase.CustomerUsecase
	validation      *validator.Validate
	wg              sync.WaitGroup
}

func NewCustomerHandlerImpl(customerUsecase usecase.CustomerUsecase) *customerHandler {
	return &customerHandler{customerUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *customerHandler) RegisterCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request customerdto.CustomerRequest
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
