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
		pubsub: NewRedisPubSub("localhost:6379"),
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.registerRoutes(),
	}
}

func handleServeWeb() http.Handler {
	return http.FileServer(http.Dir("web"))
}
