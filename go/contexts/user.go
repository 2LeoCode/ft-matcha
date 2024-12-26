package contexts

import (
	"context"
	"matcha/go/database/sqlc"
)

var userKey contextKey = "user"

func GetUser(ctx context.Context) *sqlc.GetUserFullRow {
	if user, ok := ctx.Value(userKey).(*sqlc.GetUserFullRow); !ok {
		return user
	}
	return nil
}

func WithUser(parent context.Context, user *sqlc.GetUserFullRow) context.Context {
	return context.WithValue(parent, userKey, user)
}
