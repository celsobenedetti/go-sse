package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleRoomSubscribe(broker MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomId := chi.URLParam(r, "roomId")
		userId := chi.URLParam(r, "userId")
		if len(userId) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "userId is required")
			return

		}

		ch, unsub, err := broker.Subscribe(roomId, userId)
		if err != nil {
			fmt.Fprintf(w, "error subscribing to room: %s %s", roomId, err)
			return
		}
		defer unsub()

		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Add("Cache-Control", "no-cache")

		for {
			select {
			case msg := <-ch:
				event := "message"
				encodeEvent(w, event, msg.Id, msg)
			case <-r.Context().Done():
				return
			}
		}
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
			fmt.Printf("roomId is required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
