package server

import (
	"fmt"
	"multiplicator/internal/config"
	"net/http"
)

type Service interface {
	GenerateMultiplicator() float64
}

type HTTPServer struct {
	service Service
	srv     http.Server
}

func NewHTTPServer(cfg *config.Config, service Service) *HTTPServer {
	h := &HTTPServer{
		service: service,
		srv: http.Server{
			Addr: fmt.Sprintf(":%d", cfg.Server.Port),
		},
	}

	http.Handle("GET /get", http.HandlerFunc(h.GenerationHandler))

	return h
}

func (s *HTTPServer) Start() error {
	return s.srv.ListenAndServe()
}
