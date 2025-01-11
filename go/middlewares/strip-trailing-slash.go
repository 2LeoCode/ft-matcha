package middlewares

import (
	"net/http"
)

func StripTrailingSlash(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if len(path) != 0 && path[len(path)-1] == '/' {
			http.Redirect(w, r, path[:len(path)-1], http.StatusMovedPermanently)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}
