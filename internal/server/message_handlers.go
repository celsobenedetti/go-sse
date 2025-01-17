package server

import (
	"fmt"
	"net/http"
)

type NewMessageReq struct {
	RoomID   string `json:"roomId"`
	SenderID string `json:"senderId"`
	Message  string `json:"message"`
}

func handlePostMessage(pubsub *RedisPubSub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg, err := decode[NewMessageReq](r)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprint(w, err)
			return

		}

		err = pubsub.Publish(Message{
			Id:       "TODO: C-19 message id should come from store",
			RoomId:   msg.RoomID,
			SenderId: msg.SenderID,
			Message:  msg.Message,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}

		fmt.Fprint(w, "OK")
	}
}
