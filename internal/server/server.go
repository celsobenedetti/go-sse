package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	port   int
	router *chi.Mux
	pubsub *RedisPubSub
}

func NewServer() *http.Server {
	s := &Server{
		port:   3000,
		router: chi.NewRouter(),
		pubsub: NewRedisPubSub(),
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.registerRoutes(),
	}
}
