package handlers

import (
	"matcha/go/globals"
	"net/http"
	"strconv"
)

var (
	allowCredentialsHeader string
	allowMethodsHeader     string
	maxAgeHeader           = strconv.Itoa(globals.MaxAge)
)

func init() {
	if globals.AllowCredentials {
		allowCredentialsHeader = "true"
	} else {
		allowCredentialsHeader = "false"
	}

	for i, method := range globals.AllowedMethods {
		if i != 0 {
			allowMethodsHeader += ", "
		}
		allowMethodsHeader += method
	}
}

var CorsPreflight = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", allowCredentialsHeader)
	w.Header().Set("Access-Control-Allow-Methods", allowMethodsHeader)
	w.Header().Set("Access-Control-Max-Age", maxAgeHeader)
})
