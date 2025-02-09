package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

// TODO: probably should receive a SessionService or smth
func handleCreateSession(sessionStore *SessionsStore) http.HandlerFunc {
	type Request struct {
		UserId   string `json:"user_id,omitempty"`
		Username string `json:"username,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 1. grab userId and username from request body
		req, err := decode[Request](r)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintln(w, err)
			return
		}

		if len(req.UserId) == 0 || len(req.Username) == 0 {
			// TODO: C-25 validate things
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintln(w, fmt.Errorf("user_id and username are required"))
			return
		}

		// 2. save to session store
		err = sessionStore.Upsert(req.UserId, req.Username)
		if err != nil {
			msg := fmt.Sprintf("failed to write to Redis: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, msg, req.UserId, req.Username)
			slog.Error(msg)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "session created: %s %s", req.UserId, req.Username)
	}
}
