package http

import (
	"context"
	"encoding/json"
	dto "m-banking/dto/result"
	"m-banking/interfaces/usecase"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type transactionHander struct {
	transactionUsecase usecase.TransactionUsecase
	wg                 sync.WaitGroup
	mu                 sync.Mutex
}

func NewTransactionHandlerImpl(transactionUsecase usecase.TransactionUsecase) *transactionHander {
	return &transactionHander{transactionUsecase, sync.WaitGroup{}, sync.Mutex{}}
}

func (h *transactionHander) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	accountNumber, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
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
		transactionResponse, err := h.transactionUsecase.GetTransactionUsecase(ctx, accountNumber)
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
