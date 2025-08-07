package root

import (
	"matcha/go/components"
	"matcha/go/pages/v2"
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"
)

type RootPage struct{}

func (RootPage) Path() (result string) {
	result = ""
	return
}

func (RootPage) Setup(_ *goji.Mux) (result http.Handler) {
	result = templ.Handler(components.Page("Matcha", body()))
	return
}

func (RootPage) Children() (result []v2.Page) {
	result = []v2.Page{}
	return
}
