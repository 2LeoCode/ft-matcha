package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"matcha/go/constants"
	"matcha/go/database/sqlc"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema/*.sql
var schemas string

var Handle *sqlc.Queries

func init() {
	ctx := context.Background()
	if db, err := sql.Open("sqlite3", fmt.Sprintf(
		"matcha.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=SSHA512&_auth_salt=%s&_fk=true",
		constants.DatabaseUser,
		constants.DatabasePassword,
		constants.DatabaseSalt,
	)); err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	} else if _, err := db.ExecContext(ctx, schemas); err != nil {
		log.Fatalf("Failed to initialize database tables: %s\n", err.Error())
	} else {
		Handle = sqlc.New(db)
	}
}
