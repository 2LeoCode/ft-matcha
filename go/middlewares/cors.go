package middlewares

import (
	"matcha/go/globals"
	"net/http"
)

func Cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", globals.AllowedOrigin)
		handler.ServeHTTP(w, r)
	})
}
