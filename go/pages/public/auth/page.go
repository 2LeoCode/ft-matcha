package auth

import (
	"maps"
	"matcha/go/pages"
	"matcha/go/pages/public/auth/login"
	"matcha/go/pages/public/auth/register"
	"net/http"

	"goji.io/v3"
)

var Page = pages.NewPage(&pages.PageOptions{
	Path: "/auth",
	Setup: func(mux *goji.Mux) http.Handler {
		return http.RedirectHandler("/auth/login", http.StatusPermanentRedirect)
	},
	Children: []*pages.Page{
		login.Page,
		register.Page,
	},
})
