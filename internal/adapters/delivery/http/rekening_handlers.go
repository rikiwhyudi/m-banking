package http

import (
	"context"
	"encoding/json"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type accountNumberHandlerImpl struct {
	accountNumberUsecase ports.AccountNumberUsecase
	validation           *validator.Validate
	wg                   sync.WaitGroup
}

func NewAccountNumberHandler(accountNumberUsecase ports.AccountNumberUsecase) *accountNumberHandlerImpl {
	return &accountNumberHandlerImpl{accountNumberUsecase, validator.New(), sync.WaitGroup{}}
}

func (h *accountNumberHandlerImpl) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Parse the accountNumber from the request.
	accountNumber, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberUsecase.GetBalanceUsecase(ctx, accountNumber)
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.AccountNumberRequest
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
		accountNumberResponse, err := h.accountNumberUsecase.DepositUsecase(ctx, request)
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberUsecase.CashoutUsecase(ctx, request)
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request dto.TransferBalanceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err = h.validation.Struct(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberUsecase.TransferUsecase(ctx, request)
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
