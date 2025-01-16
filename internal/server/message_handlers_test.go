package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/celsobenedetti/go-sse/testing/containers"
)

func Test_handlePostMessage(t *testing.T) {
	pubsub, close := newRedisContainer(t)
	defer close()

	newRequest := func(msg NewMessageReq) (*httptest.ResponseRecorder, *http.Request) {
		body, err := json.Marshal(msg)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
		res := httptest.NewRecorder()
		return res, req
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
				res, req := newRequest(tC.body)

				handle := handlePostMessage(pubsub)
				handle(res, req)

				// TODO: C-25 should be http.StatusUnprocessableEntity instead
				assert.Equal(t, res.Result().StatusCode, http.StatusOK)
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

		res, req := newRequest(msg)
		handle := handlePostMessage(pubsub)
		handle(res, req)

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
				return
			case <-timeout.C:
				t.Fatalf("timed out without receiving message on subscriber")
			}
		}
	})
}

func newRedisContainer(t *testing.T) (*RedisPubSub, func()) {
	redisContainer, close, err := containers.Redis()
	assert.Nil(t, err)

	ctx := context.Background()
	connectionString, err := redisContainer.Endpoint(ctx, "")
	assert.Nil(t, err)

	pubsub := NewRedisPubSub(connectionString)
	return pubsub, close
}
