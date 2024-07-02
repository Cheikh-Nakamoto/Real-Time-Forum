package handlers

import (
	"net/http"
)

func Action(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/action" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
