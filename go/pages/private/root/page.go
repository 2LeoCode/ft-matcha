package root

import (
	"matcha/go/components"
	"matcha/go/pages"
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"
)

var Page = pages.NewPage(func(path string, mux *goji.Mux) http.Handler {
	return templ.Handler(components.Page("Matcha", body()))
})
