package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return nil
}

func decode[T any](r *http.Request) (v T, err error) {
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
