package validator

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ValidationResponse struct {
	Valid bool           `json:"valid"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidateHandler struct {
	validator *CardValidator
}

func NewValidateHandler(v *CardValidator) *ValidateHandler {
	return &ValidateHandler{validator: v}
}

func (h *ValidateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var c Card
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		internalServerError(w)
		return
	}

	err := h.validator.Validate(&c)
	switch {
	case errors.Is(err, ErrCardExpired):
		badRequestError(w, ErrCodeCardExpired, err.Error())
		return
	case errors.Is(err, ErrInvalidCardNumberLength):
		badRequestError(w, ErrCodeInvalidCardNumberLength, err.Error())
		return
	case errors.Is(err, ErrInvalidMonth):
		badRequestError(w, ErrCodeInvalidMonth, err.Error())
		return
	}
	if err != nil {
		log.Println(err)
		internalServerError(w)
		return
	}

	json.NewEncoder(w).Encode(&ValidationResponse{
		Valid: true,
	})
}

func badRequestError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&ValidationResponse{
		Valid: false,
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
		},
	})
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&ValidationResponse{
		Valid: false,
		Error: &ErrorResponse{
			Code:    ErrCodeInternal,
			Message: "Internal Server Error",
		},
	})
}
