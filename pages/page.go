package pages

import "goji.io/v3"

type Page interface {
	Register(root *goji.Mux)
}
