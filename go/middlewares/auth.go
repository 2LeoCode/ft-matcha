package middlewares

import (
	"context"
	"database/sql"
	"log"
	"matcha/go/contexts"
	"matcha/go/database"
	"matcha/go/globals"
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

type tokenData struct {
	mail     string
	password string
}

func verifyUser(claims jwt.MapClaims) *sqlc.GetUserFullRow {
	data := claims["data"].(tokenData)
	hasher := sha3.New512()

	hasher.Write([]byte(data.password))
	param := sqlc.GetUserFullParams{
		Mail:     data.mail,
		Password: hasher.Sum(nil),
	}

	if user, err := database.Handle.GetUserFull(context.Background(), param); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatalf("Error while fetching user from database: %s\n", err.Error())
	} else {
		return &user
	}
	return nil
}

func Auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *sqlc.GetUserFullRow
		defer func() {
			if user != nil {
				handler.ServeHTTP(w, r.WithContext(contexts.WithUser(r.Context(), user)))
			} else {
				http.Redirect(w, r, "/auth", http.StatusPermanentRedirect)
			}
		}()

		matchaCookies := r.CookiesNamed("matcha_token")
		if i := slices.IndexFunc(matchaCookies, func(c *http.Cookie) bool {
			return c.Domain == globals.ServerHost && c.Path == "/"
		}); i == -1 {
			return
		} else if token, err := jwt.Parse(matchaCookies[i].Value, func(token *jwt.Token) (interface{}, error) {
			// For now we use a hard-coded key with HMAC/SHA512 for debug,
			// but in production we will use ECDSA with a public/private
			// key pair and a secret passphrase stored in files.
			// We also don't use the expiration claim of the token,
			// but in production tokens will have an expiration time.
			return globals.JwtSecret, nil
		}, jwt.WithValidMethods([]string{"HS512"})); err == nil {
			user = verifyUser(token.Claims.(jwt.MapClaims))
		}
	})
}
