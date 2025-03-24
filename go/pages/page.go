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
	options := this.options
	if options == nil {
		log.Fatalln("Cannot register a page twice")
	}
	submux := goji.SubMux()
	for _, page := range options.Children {
		page.Register(submux)
	}
	mux.Handle(pat.Get(options.Path), options.Setup(submux))
	mux.Handle(pat.NewWithMethods(fmt.Sprintf("%s/*", options.Path), http.MethodGet, http.MethodPost), submux)
	this.options = nil
}
