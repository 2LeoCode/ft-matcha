package main

import (
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func main() {
	root := goji.NewMux()

	// Serve static CSS and JS files
	root.Handle(pat.Get("/css/*"), http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	root.Handle(pat.Get("/js/*"), http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	// Register page handlers
	root.Handle(pat.Get("/"), templ)
}
