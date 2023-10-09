package handlers

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	}
}
