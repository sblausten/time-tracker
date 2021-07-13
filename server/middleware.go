package server

import (
	"log"
	"net/http"
)

func logRequest(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recieved request to %s %s from %s", r.Method, r.URL.Path, r.Referer())
		f(w, r)
	}
}
