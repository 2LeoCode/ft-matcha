package v2

import (
	"fmt"
	"net/http"

	"goji.io/v3"
	"goji.io/v3/pat"
)

type Page interface {
	Path() string
	Setup(mux *goji.Mux) http.Handler
	Children() []Page
}

type router struct {
	mux *goji.Mux
}

func NewRouter(mux *goji.Mux) (result *router) {
	result = &router{mux}
	return
}

func (this *router) Register(page Page) {
	submux := goji.SubMux()
	for _, page := range page.Children() {
		this.Register(page)
	}
	this.mux.Handle(pat.Get(page.Path()), page.Setup(submux))
	this.mux.Handle(pat.NewWithMethods(fmt.Sprintf("%s/*", page.Path()), http.MethodGet, http.MethodPost), submux)
}
