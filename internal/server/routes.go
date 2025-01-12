package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Handle("/web/*", http.StripPrefix("/web", handleServeWeb()))
	r.Get("/healthz", handleHealthz())

	r.Route("/rooms", func(r chi.Router) {
		r.With(checkRoomId).Route("/{roomId}", func(r chi.Router) {
			r.Get("/messages", handleGetRoomMessages())
			r.Get("/subscribe/{userId}", handleRoomSubscribe(s.messageBroker))
		})
	})

	r.Route("/messages", func(r chi.Router) {
		r.Post("/", handlePostMessage(s.messageBroker))
	})

	return r
}
