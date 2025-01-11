package auth

import (
	"matcha/go/pages/auth/login"
	"matcha/go/pages/auth/register"
	"net/http"

	"goji.io/v3"
	"goji.io/v3/pat"
)

func Register(mux *goji.Mux) {
	mux.Handle(pat.Get("/auth"), http.RedirectHandler("/auth/login", http.StatusMovedPermanently))

	subMux := goji.SubMux()
	login.Register(subMux)
	register.Register(subMux)
	mux.Handle(pat.NewWithMethods("/auth/*", http.MethodGet, http.MethodPost), subMux)
}
