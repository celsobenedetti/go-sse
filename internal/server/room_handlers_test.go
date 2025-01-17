package server

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/celsobenedetti/go-sse/testing/containers"
)

func Test_handleRoomSubscribe(t *testing.T) {
	pubsub, close := runRedisTestContainer(t)
	defer close()

	handle := handleRoomSubscribe(pubsub)

	t.Run("should receive published message", func(t *testing.T) {
		roomId := "123"
		userId := "456"
		msg := Message{
			Id:       "id",
			RoomId:   roomId,
			SenderId: userId,
			Message:  "hello from test",
		}

		req := httptest.NewRequest("GET", "/rooms/"+roomId+"/users/"+userId, nil)
		w := httptest.NewRecorder()

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("roomId", roomId)
		rctx.URLParams.Add("userId", userId)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		go handle(w, req)

		// Give some time for the handler to subscribe
		time.Sleep(100 * time.Millisecond)

		// Publish the message
		err := pubsub.Publish(msg)
		assert.Nil(t, err)

		// Give some time for the message to be processed
		time.Sleep(100 * time.Millisecond)

		msgJSON, err := msg.MarshalBinary()
		assert.Nil(t, err)

		// Check the response
		response := w.Body.String()
		assert.Contains(t, response, "id: id")
		assert.Contains(t, response, "event: message")
		assert.Contains(t,
			normalize(response),
			normalize(fmt.Sprintf("data: %s", string(msgJSON))),
		)
	})
}

// Normalize string to account for characters included in JSON encoding
func normalize(s string) string {
	s = strings.ReplaceAll(s, `"`, "")
	s = strings.ReplaceAll(s, "\\", "")
	return s
}

func runRedisTestContainer(t *testing.T) (*RedisPubSub, func()) {
	redisContainer, close, err := containers.Redis()
	assert.Nil(t, err)

	ctx := context.Background()
	connectionString, err := redisContainer.Endpoint(ctx, "")
	assert.Nil(t, err)

	pubsub := NewRedisPubSub(connectionString)
	return pubsub, close
}
