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
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		writeResponse(w, http.StatusBadRequest, response)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.GetBalanceService(accountNumber)
		if err != nil {
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			writeResponse(w, http.StatusBadRequest, response)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		writeResponse(w, http.StatusOK, response)
	}()

	h.wg.Wait()
}

func (h *accountNumberHandlerImpl) DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		writeResponse(w, http.StatusBadRequest, response)
		return
	}

	// Validate request input using go-playground/validator
	if err = h.validation.Struct(request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		writeResponse(w, http.StatusBadRequest, response)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.DepositService(request)
		if err != nil {
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			writeResponse(w, http.StatusBadRequest, response)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		writeResponse(w, http.StatusOK, response)
	}()

	h.wg.Wait()
}

func (h *accountNumberHandlerImpl) CashoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var request accNumberdto.AccountNumberRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		writeResponse(w, http.StatusBadRequest, response)
		return
	}

	if err = h.validation.Struct(request); err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		writeResponse(w, http.StatusBadRequest, response)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		accountNumberResponse, err := h.accountNumberServiceImpl.CashoutService(request)
		if err != nil {
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			writeResponse(w, http.StatusBadRequest, response)
			return
		}

		response := dto.SuccessResult{Code: http.StatusOK, Data: accountNumberResponse}
		writeResponse(w, http.StatusOK, response)
	}()

	h.wg.Wait()
}
