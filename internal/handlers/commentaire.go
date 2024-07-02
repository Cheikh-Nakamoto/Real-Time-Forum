package handlers

import (
	"net/http"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && r.URL.Path == "/comment" {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
    }
}
