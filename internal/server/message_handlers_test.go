package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/celsobenedetti/go-sse/testing/containers"
)

func Test_handlePostMessage(t *testing.T) {
	t.Run("should return 422 if body does not include valid message", func(t *testing.T) {
		pubsub, close := newRedisContainer(t)
		defer close()

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
				body, err := json.Marshal(tC.body)
				assert.Nil(t, err)

				req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))
				res := httptest.NewRecorder()

				handle := handlePostMessage(pubsub)
				handle(res, req)

				assert.Equal(t, res.Result().StatusCode, http.StatusUnprocessableEntity)
				assert.Equal(t, "true", "true")
			})
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
