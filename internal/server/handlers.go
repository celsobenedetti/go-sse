package server

import "net/http"

func handleServeWeb() http.Handler {
	return http.FileServer(http.Dir("web"))
}

func handleHealthz(redis *RedisPubSub) http.HandlerFunc {
	errors := []error{}

	checks := map[string]string{
		"messageStore": "TODO",
	}

	if msg, err := redis.Health(); err != nil {
		checks["Redis"] = err.Error()
		errors = append(errors, err)
	} else {
		checks["Redis"] = msg
	}

	status := http.StatusOK
	if len(errors) > 0 {
		status = http.StatusInternalServerError
	}

	return func(w http.ResponseWriter, r *http.Request) {
		encode(w, checks, status)
	}
}
