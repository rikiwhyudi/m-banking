package handlers

import (
	"context"
	"encoding/json"
	accNumberdto "m-banking/dto/rekening"
	dto "m-banking/dto/result"
	"m-banking/service"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type accountNumberHandlerImpl struct {
	accountNumberServiceImpl service.AccountNumberService
	validation               *validator.Validate
	wg                       sync.WaitGroup
	mu                       sync.Mutex
}

func NewHandlerAccountNumberImpl(accountNumberServiceImpl service.AccountNumberService) *accountNumberHandlerImpl {
	return &accountNumberHandlerImpl{accountNumberServiceImpl, validator.New(), sync.WaitGroup{}, sync.Mutex{}}
}

func (h *accountNumberHandlerImpl) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Parse the accountNumber from the request
	accountNumber, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	h.mu.Lock()
	defer h.mu.Unlock()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.GetBalanceService(ctx, accountNumber)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}

func (h *accountNumberHandlerImpl) DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
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
		accountNumberResponse, err := h.accountNumberServiceImpl.DepositService(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}

func (h *accountNumberHandlerImpl) CashoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

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
		accountNumberResponse, err := h.accountNumberServiceImpl.CashoutService(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}

func (h *accountNumberHandlerImpl) TransferHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.TransferBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

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
		accountNumberResponse, err := h.accountNumberServiceImpl.TransferService(ctx, request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}
