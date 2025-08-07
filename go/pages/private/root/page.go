package root

import (
	"matcha/go/components"
	"matcha/go/pages"
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"
)

var Page = pages.NewPage(&pages.PageOptions{
	Path: "",
	Setup: func(mux *goji.Mux) (root http.Handler) {
		root = templ.Handler(components.Page("Matcha", body(), true))
		return
	},
})
