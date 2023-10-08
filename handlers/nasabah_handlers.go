package handlers

import (
	customerdto "e-wallet/dto/nasabah"
	dto "e-wallet/dto/result"
	"e-wallet/service"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type customerHandlerImpl struct {
	customerServiceImpl service.CustomerService
	validation          *validator.Validate
}

func NewHanlderCustomerImpl(customerServiceImpl service.CustomerService) *customerHandlerImpl {
	return &customerHandlerImpl{customerServiceImpl, validator.New()}
}

func (h *customerHandlerImpl) RegisterCustomerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request customerdto.CustomerRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewDecoder(r.Body).Decode(&response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	customerResponse, err := h.customerServiceImpl.RegisterCustomerService(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.Result{Code: http.StatusOK, Message: customerResponse}
	json.NewEncoder(w).Encode(response)

}
