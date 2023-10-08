package handlers

import (
	accNumberdto "e-wallet/dto/rekening"
	dto "e-wallet/dto/result"
	"e-wallet/service"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type accountNumberHandlerImpl struct {
	accountNumberServiceImpl service.AccountNumberService
	validation               *validator.Validate
}

func NewHandlerAccountNumberImpl(accountNumberServiceImpl service.AccountNumberService) *accountNumberHandlerImpl {
	return &accountNumberHandlerImpl{accountNumberServiceImpl, validator.New()}
}

func (h *accountNumberHandlerImpl) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	accountNumber, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	responses := make(chan dto.Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.GetBalanceService(accountNumber)
		if err != nil {
			responses <- dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
			return
		}
		responses <- dto.Result{Code: http.StatusOK, Message: accountNumberResponse}
	}()

	wg.Wait()
	close(responses)

	for response := range responses {
		if response.Code != http.StatusOK {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (h *accountNumberHandlerImpl) DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	responses := make(chan dto.Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.DepositService(request)
		if err != nil {
			responses <- dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
			return
		}
		responses <- dto.Result{Code: http.StatusOK, Message: accountNumberResponse}
	}()

	wg.Wait()
	close(responses)

	for response := range responses {
		if response.Code != http.StatusOK {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (h *accountNumberHandlerImpl) CashoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	responses := make(chan dto.Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.CashoutService(request)
		if err != nil {
			responses <- dto.Result{Code: http.StatusBadRequest, Message: err.Error()}
			return
		}
		responses <- dto.Result{Code: http.StatusOK, Message: accountNumberResponse}
	}()

	wg.Wait()
	close(responses)

	for response := range responses {
		if response.Code != http.StatusOK {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
