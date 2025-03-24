package login

import (
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"

	"matcha/go/components"
	"matcha/go/pages"
)

var Page = pages.NewPage(&pages.PageOptions{
	Path: "login",
	Setup: func(mux *goji.Mux) http.Handler {
		return templ.Handler(components.Page("Login", body()))
	},
})
