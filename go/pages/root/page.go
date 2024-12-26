package root

import (
	"matcha/go/components"

	"github.com/a-h/templ"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func Register(mux *goji.Mux) {
	mux.Handle(pat.Get(""), templ.Handler(components.Page("Matcha", body(), true)))
}
