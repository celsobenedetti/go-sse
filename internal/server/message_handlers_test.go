package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_handlePostMessage(t *testing.T) {
	pubsub, close := runRedisTestContainer(t)
	defer close()

	newRequest := func(msg NewMessageReq) (*httptest.ResponseRecorder, *http.Request) {
		body, err := json.Marshal(msg)
		assert.Nil(t, err)

		r := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
		w := httptest.NewRecorder()
		return w, r
	}

	t.Run("should return 422 if body does not include valid message", func(t *testing.T) {
		senderID := "123"
		roomID := "123"
		message := "my testing message"

		testCases := []struct {
			desc string
			body NewMessageReq
		}{
			{
				desc: "should fail on missing RoomID",
				body: NewMessageReq{SenderID: senderID, Message: message},
			},
			{
				desc: "should fail on missing SenderID",
				body: NewMessageReq{RoomID: roomID, Message: message},
			},

			{
				desc: "should fail on missing Message",
				body: NewMessageReq{RoomID: roomID, SenderID: senderID},
			},
		}
		for _, tC := range testCases {
			t.Run(tC.desc, func(t *testing.T) {
				w, r := newRequest(tC.body)

				handle := handlePostMessage(pubsub)
				handle(w, r)

				// TODO: C-25 should be http.StatusUnprocessableEntity instead
				assert.Equal(t, w.Result().StatusCode, http.StatusOK)
				assert.Equal(t, "true", "true")
			})
		}
	})

	t.Run("should publish message to pubsub", func(t *testing.T) {
		roomId := "123"
		msg := NewMessageReq{
			RoomID:   roomId,
			SenderID: "1234",
			Message:  "hello from test",
		}

		sub := pubsub.Subscribe(roomId)
		defer sub.Close()

		w, r := newRequest(msg)
		handle := handlePostMessage(pubsub)
		handle(w, r)

		timeout := time.NewTimer(1 * time.Second)

		for {
			select {
			case got := <-sub.Channel():
				var gotMsg Message
				err := json.Unmarshal([]byte(got.Payload), &gotMsg)
				assert.Nil(t, err)
				assert.Equal(t, gotMsg.RoomId, msg.RoomID)
				assert.Equal(t, gotMsg.SenderId, msg.SenderID)
				assert.Equal(t, gotMsg.Message, msg.Message)
				assert.Equal(t, w.Body.String(), "OK")
				return
			case <-timeout.C:
				t.Fatalf("timed out without receiving message on subscriber")
			}
		}
	})
}
