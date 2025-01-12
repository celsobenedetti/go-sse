package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	port          int
	router        *chi.Mux
	messageBroker MessageBroker
}

func NewServer() *http.Server {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	default:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	s := &Server{
		port:          3000,
		router:        chi.NewRouter(),
		messageBroker: NewInMemoryMessageBroker(),
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.registerRoutes(),
	}
}
