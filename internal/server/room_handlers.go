package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleRoomSubscribe(pubsub *RedisPubSub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomId := chi.URLParam(r, "roomId")
		// TODO: should come from headers and middleware
		userId := chi.URLParam(r, "userId")
		if len(userId) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "userId is required")
			return

		}

		sub := pubsub.Subscribe(roomId)
		defer sub.Close()

		w.Header().Add("Content-Type", "text/event-stream")
		w.Header().Add("Cache-Control", "no-cache")

		for {
			select {
			case msg := <-sub.Channel():
				event := "message"
				err := encodeEvent(w, event, "id", msg.Payload)
				if err != nil {
					fmt.Fprintf(w, "error encoding event in event-stream: %s", err.Error())
					return
				}
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
