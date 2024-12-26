package main

import (
	"fmt"
	"matcha/go/components"
	"matcha/go/globals"
	"matcha/go/middlewares"
	"matcha/go/pages/auth"
	"matcha/go/pages/root"
	"net/http"

	"github.com/a-h/templ"
	"goji.io/v3"
	"goji.io/v3/pat"
)

func main() {
	mux := goji.NewMux()
	// Contains all pages that require auth
	protectedMux := goji.SubMux()

	mux.Use(middlewares.StripTrailingSlash)
	protectedMux.Use(middlewares.Auth)

	// 404 Not Found fallback
	mux.Handle(pat.Get("*"), templ.Handler(components.Page("Not found", components.NotFound(), false)))

	// Static CSS and JS files
	mux.Handle(pat.Get("/css/*"), http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle(pat.Get("/js/*"), http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	mux.Handle(pat.NewWithMethods("/*", http.MethodGet, http.MethodPost), protectedMux)

	// Register page handlers
	auth.Register(mux)
	root.Register(protectedMux)

	// TODO: use https instead of http for production
	http.ListenAndServe(fmt.Sprintf("%s:%d", globals.ServerHost, globals.ServerPort), mux)
}
