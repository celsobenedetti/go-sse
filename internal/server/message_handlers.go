package server

import (
	"fmt"
	"net/http"
)

func handlePostMessage(broker MessageBroker) http.HandlerFunc {
	type NewMessageReq struct {
		RoomId   string `json:"roomId"`
		SenderId string `json:"senderId"`
		Message  string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		msg, err := decode[NewMessageReq](r)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprint(w, err)
			return

		}

		err = broker.Publish(Message{
			Id:       "TODO: we don't have message id",
			RoomId:   msg.RoomId,
			SenderId: msg.SenderId,
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
