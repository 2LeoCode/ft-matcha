package auth

import (
	"matcha/go/pages/auth/login"
	"matcha/go/pages/auth/register"
	"net/http"

	"goji.io/v3"
	"goji.io/v3/pat"
)

func Register(mux *goji.Mux) {
	mux.HandleFunc(pat.Get("/auth"), func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/auth/login", http.StatusMovedPermanently)
	})

	subMux := goji.SubMux()
	login.Register(subMux)
	register.Register(subMux)
	mux.Handle(pat.NewWithMethods("/auth/*", http.MethodGet, http.MethodPost), subMux)
}
