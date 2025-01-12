package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleServeWeb() http.Handler {
	return http.FileServer(http.Dir("web"))
}

func handleHealthz(messageBroker MessageBroker) http.HandlerFunc {
	errors := []error{}

	checks := map[string]string{
		"messageStore": "TODO",
	}

	if msg, err := messageBroker.Health(); err != nil {
		checks["messageBroker"] = err.Error()
		errors = append(errors, err)
	} else {
		checks["messageBroker"] = msg
	}

	status := http.StatusOK
	if len(errors) > 0 {
		status = http.StatusInternalServerError
	}

	return func(w http.ResponseWriter, r *http.Request) {
		encode(w, checks, status)
	}
}

func encodeEvent[T any](w http.ResponseWriter, event, id string, data T) error {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return fmt.Errorf("Streaming unsupported")
	}

	buf, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error encodign event json data: %w", err)
	}

	w.Write([]byte(fmt.Sprintf("id: %s\n", id)))
	w.Write([]byte(fmt.Sprintf("event: %s\n", event)))
	w.Write([]byte(fmt.Sprintf("data: %s\n", string(buf))))
	w.Write([]byte("\n\n"))
	flusher.Flush()
	return nil
}

func encode[T any](w http.ResponseWriter, v T, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (v T, err error) {
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
