package utils

import (
	"encoding/json"
	"net/http"
)

type HandelApiErrorType func(http.ResponseWriter, *http.Request) error

func HandleApiError(f HandelApiErrorType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respErr := map[string]string{"error": err.Error()}
			if jsonErr := json.NewEncoder(w).Encode(respErr); jsonErr != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
		}
	}
}
