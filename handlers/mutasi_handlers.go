package handlers

import (
	dto "e-wallet/dto/result"
	"e-wallet/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type transactionHandlerImpl struct {
	transactionServiceImpl service.TransactionService
}

func NewHandlerTransactionImpl(transactionServiceImpl service.TransactionService) *transactionHandlerImpl {
	return &transactionHandlerImpl{transactionServiceImpl}
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

	transactionResponse, err := h.transactionServiceImpl.GetTransactionService(accountNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transactionResponse}
	json.NewEncoder(w).Encode(response)
}
