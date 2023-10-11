package handlers

import (
	"context"
	customerdto "e-wallet/dto/nasabah"
	dto "e-wallet/dto/result"
	"e-wallet/service"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type customerHandlerImpl struct {
	customerServiceImpl service.CustomerService
	validation          *validator.Validate
	wg                  sync.WaitGroup
	mu                  sync.Mutex
}

func NewHanlderCustomerImpl(customerServiceImpl service.CustomerService) *customerHandlerImpl {
	return &customerHandlerImpl{customerServiceImpl, validator.New(), sync.WaitGroup{}, sync.Mutex{}}
}

func (h *customerHandlerImpl) RegisterCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request customerdto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewDecoder(r.Body).Decode(&response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		customerResponse, err := h.customerServiceImpl.RegisterCustomerService(ctx, request)
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
