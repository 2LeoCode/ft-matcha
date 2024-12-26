package middlewares

import (
	"context"
	"log"
	"matcha/go/contexts"
	"matcha/go/database"
	"matcha/go/database/sqlc"
	"matcha/go/globals"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func verifyUser(claims jwt.MapClaims) *sqlc.GetUserFullRow {
	tokenData := claims["data"].(sqlc.GetUserFullParams)

	// fetch user in database and
	user, err := database.Handle.GetUserFull(context.Background(), tokenData)
	if err != nil {
		log.Fatalf("Error while fetching user from database: %s\n", err.Error())
	}
	return &user
}

func Auth(handler http.Handler) http.Handler {
	mw := func(w http.ResponseWriter, r *http.Request) {
		var user *sqlc.GetUserFullRow
		defer func() {
			if user != nil {
				handler.ServeHTTP(w, r.WithContext(contexts.WithUser(r.Context(), user)))
			} else {
				http.Redirect(w, r, "/auth", http.StatusPermanentRedirect)
			}
		}()

		authHeader := r.Header.Get("Authentication")
		if authHeader[:7] != "Bearer " {
			return
		}
		if token, err := jwt.Parse(authHeader[7:], func(token *jwt.Token) (interface{}, error) {
			// For now we use a hard-coded key with HMAC/SHA512 for debug,
			// but in production we will use ECDSA with a public/private
			// key pair and a secret passphrase stored in files.
			// We also don't use the expiration claim of the token,
			// but in production tokens will have an expiration time.
			return globals.JwtSecret, nil
		}, jwt.WithValidMethods([]string{"HS512"})); err == nil {
			user = verifyUser(token.Claims.(jwt.MapClaims))
		}
	}
	return http.HandlerFunc(mw)
}
