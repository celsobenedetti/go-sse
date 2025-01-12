package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleRoomSubscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func handleGetRoomMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "handleGetRoomMessages")
	}
}

// Write 400 BadRequest if roomId Param is empty
func checkRoomId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roomId := chi.URLParam(r, "roomId")
		if len(roomId) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "roomId is required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
