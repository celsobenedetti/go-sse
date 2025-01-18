package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		// TODO: allow specific domains
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Handle("/web/*", http.StripPrefix("/web", handleServeWeb()))
	r.Get("/healthz", handleHealthz(s.pubsub))

	r.Route("/rooms", func(r chi.Router) {
		r.With(checkRoomId).Route("/{roomId}", func(r chi.Router) {
			r.Get("/messages", handleGetRoomMessages())
			r.Get("/subscribe/{userId}", handleRoomSubscribe(s.pubsub))
		})
	})

	r.Route("/messages", func(r chi.Router) {
		r.Post("/", handlePostMessage(s.pubsub))
	})

	return r
}
