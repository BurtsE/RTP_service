package server

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidValue    = errors.New("invalid value")
	ErrValueOutOfRange = errors.New("value out of range")
)

type ErrorResponceDTO struct {
	Error string `json:"error"`
}

func writeErrorResponce(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponceDTO{Error: err.Error()})

}
