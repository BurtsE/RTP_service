package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type GenerationResponseDTO struct {
	Multiplicator float64 `json:"result"`
}

func (s *HTTPServer) GenerationHandler(w http.ResponseWriter, r *http.Request) {
	multiplicator := s.service.GenerateMultiplicator()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerationResponseDTO{Multiplicator: multiplicator})

	log.Printf("generated multiplicator: %f", multiplicator)
}
