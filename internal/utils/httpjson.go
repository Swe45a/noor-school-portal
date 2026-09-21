package utils

import (
	"encoding/json"
	"net/http"

	apperrors "school-portal/internal/errors"
)

func WriteJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func DecodeJSON(r *http.Request, dst interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return apperrors.ErrInvalidRequest
	}
	return nil
}

// WriteError maps a sentinel error to its HTTP status and snake_case wire code.
func WriteError(w http.ResponseWriter, err error) {
	status := apperrors.StatusFor(err)
	code := apperrors.Code(err)
	WriteJSON(w, status, map[string]string{"error": code})
}
