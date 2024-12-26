package login

import (
	"github.com/a-h/templ"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func Register(mux *goji.Mux) {
	mux.Handle(pat.Get("/login"), templ.Handler(body()))
}
