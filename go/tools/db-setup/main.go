package main

import (
	_ "embed"
	"log"

	"matcha/go/database"
)

//go:embed schema/tables.sql
var tablesStmt string

//go:embed schema/indexes.sql
var indexesStmt string

//go:embed schema/triggers.sql
var triggersStmt string

var schemaStmt = tablesStmt

func main() {
	if db, err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	} else if _, err := db.Exec(schemaStmt); err != nil {
		log.Fatalf("Failed to initialize database tables: %s\n", err.Error())
	}
}
