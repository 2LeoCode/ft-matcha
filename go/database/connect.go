package database

import (
	"database/sql"
	"fmt"
	"matcha/go/globals"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	return sql.Open("sqlite3", fmt.Sprintf(
		"matcha.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=SSHA512&_auth_salt=%s&_fk=true",
		globals.DatabaseUser,
		globals.DatabasePassword,
		globals.DatabaseSalt,
	))
}
