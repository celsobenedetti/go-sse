package server

import (
	"fmt"
	"net/http"
)

func handleCreateSession() http.HandlerFunc {
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

		// 2. save to redis KV
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "saved to KV: %s %s", req.UserId, req.Username)
	}
}
