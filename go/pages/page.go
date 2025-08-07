package pages

import (
	"fmt"
	"log"
	"net/http"

	"goji.io/v3"
	"goji.io/v3/pat"
)

type PageSetup func(mux *goji.Mux) (root http.Handler)

type PageOptions struct {
	Path     string
	Setup    PageSetup
	Children []*Page
}

type Page struct {
	options *PageOptions
}

func NewPage(options *PageOptions) *Page {
	if options.Setup == nil {
		log.Fatalln("Page setup function cannot be nil")
	}
	return &Page{options}
}

func (this *Page) Register(mux *goji.Mux) {
}
