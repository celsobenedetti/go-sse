package server

import (
	"fmt"
	"net/http"
)

func handleHealthz(redis *RedisPubSub) http.HandlerFunc {
	checks := map[string]string{}
	errors := []error{}

	errors = checkRedis(redis, checks, errors)

	status := http.StatusOK
	if len(errors) > 0 {
		status = http.StatusInternalServerError
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := encode(w, checks, status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error encoding healthz payload: %s", err.Error())
		}
	}
}

func checkRedis(redis *RedisPubSub, checks map[string]string, errors []error) []error {
	m, err := redis.Health()
	if err != nil {
		errors = append(errors, err)
		m = err.Error()
	}
	checks["Redis"] = m
	return errors
}
