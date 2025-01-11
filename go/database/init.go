package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"matcha/go/database/sqlc"
	"matcha/go/globals"
	"matcha/go/utils"
	"os"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var queries *sqlc.Queries

func loadSqlSchema(directory string) string {
	var schema string
	for _, entry := range utils.Must(os.ReadDir(directory)) {
		fullPath := path.Join(directory, entry.Name())
		if entry.IsDir() {
			schema += loadSqlSchema(fullPath)
		} else {
			if !strings.HasSuffix(entry.Name(), ".sql") {
				log.Fatalf("Invalid sql file name: %s\n", entry.Name())
			}
			schema += string(utils.Must(os.ReadFile(fullPath)))
		}
	}
	return schema
}

func init() {
	ctx := context.Background()
	if db, err := sql.Open("sqlite3", fmt.Sprintf(
		"matcha.db?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=SSHA512&_auth_salt=%s&_fk=true",
		globals.DatabaseUser,
		globals.DatabasePassword,
		globals.DatabaseSalt,
	)); err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	} else {
		schema := loadSqlSchema("sql/schema")
		if _, err := db.ExecContext(ctx, schema); err != nil {
			log.Fatalf("Failed to initialize database tables: %s\n", err.Error())
		} else {
			queries = sqlc.New(db)
		}
	}
}
