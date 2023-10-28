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

	"github.com/gorilla/mux"
)

type transactionHanderImpl struct {
	transactionUsecase ports.TransactionUsecase
	wg                 sync.WaitGroup
}

func NewTransactionHandler(transactionUsecase ports.TransactionUsecase) *transactionHanderImpl {
	return &transactionHanderImpl{transactionUsecase, sync.WaitGroup{}}
}

func (h *transactionHanderImpl) FindTransactionHandler(w http.ResponseWriter, r *http.Request) {
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
		transactionResponse, err := h.transactionUsecase.FindTransactionUsecase(ctx, accountNumber)
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
