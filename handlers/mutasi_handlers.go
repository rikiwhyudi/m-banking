package handlers

import (
	"context"
	dto "e-wallet/dto/result"
	"e-wallet/service"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type transactionHandlerImpl struct {
	transactionServiceImpl service.TransactionService
	wg                     sync.WaitGroup
	mu                     sync.Mutex
}

func NewHandlerTransactionImpl(transactionServiceImpl service.TransactionService) *transactionHandlerImpl {
	return &transactionHandlerImpl{transactionServiceImpl, sync.WaitGroup{}, sync.Mutex{}}
}

func (h *transactionHandlerImpl) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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
		transactionResponse, err := h.transactionServiceImpl.GetTransactionService(ctx, accountNumber)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: http.StatusOK, Data: transactionResponse}
		json.NewEncoder(w).Encode(response)
	}()

	h.wg.Wait()
}
