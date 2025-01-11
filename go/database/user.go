package database

import "matcha/go/database/sql"

type user struct{}

type userError error

type userErrors struct {
	RegisterDuplicateMail
}

var User = &user{}

var ()

func (*user) Register(mail string, password string) error {
	return sqlc.User_Create()
}
