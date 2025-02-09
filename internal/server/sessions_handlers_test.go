package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleCreateSession(t *testing.T) {
	t.Run("should succeed if user_id and username are present in body JSON", func(t *testing.T) {
		type Request struct {
			UserID   string `json:"user_id,omitempty"`
			Username string `json:"username,omitempty"`
		}

		body, err := json.Marshal(Request{
			UserID:   "fa367e3c-3bbd-4f2c-a7dc-10a251cfac09",
			Username: "John Tester",
		})
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewReader(body))

		handle := handleCreateSession()
		handle(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	})

	t.Run("should fail if user_id is not presetnt in body", func(t *testing.T) {
		type Request struct {
			Username string `json:"username,omitempty"`
		}

		body, err := json.Marshal(Request{Username: "John Tester"})
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewReader(body))

		handle := handleCreateSession()
		handle(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusUnprocessableEntity)
	})

	t.Run("should fail if username is not presetnt in body", func(t *testing.T) {
		type Request struct {
			UserID string `json:"user_id,omitempty"`
		}

		body, err := json.Marshal(Request{UserID: "7bc7f7c5-c05f-4c43-ad3f-bd1ef0b644db"})
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewReader(body))

		handle := handleCreateSession()
		handle(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusUnprocessableEntity)
	})
}
