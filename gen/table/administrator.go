//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Administrator = newAdministratorTable("", "administrator", "")

type administratorTable struct {
	sqlite.Table

	// Columns
	ID       sqlite.ColumnInteger
	Username sqlite.ColumnString
	Password sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type AdministratorTable struct {
	administratorTable

	EXCLUDED administratorTable
}

// AS creates new AdministratorTable with assigned alias
func (a AdministratorTable) AS(alias string) *AdministratorTable {
	return newAdministratorTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AdministratorTable with assigned schema name
func (a AdministratorTable) FromSchema(schemaName string) *AdministratorTable {
	return newAdministratorTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AdministratorTable with assigned table prefix
func (a AdministratorTable) WithPrefix(prefix string) *AdministratorTable {
	return newAdministratorTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AdministratorTable with assigned table suffix
func (a AdministratorTable) WithSuffix(suffix string) *AdministratorTable {
	return newAdministratorTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAdministratorTable(schemaName, tableName, alias string) *AdministratorTable {
	return &AdministratorTable{
		administratorTable: newAdministratorTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newAdministratorTableImpl("", "excluded", ""),
	}
}

func newAdministratorTableImpl(schemaName, tableName, alias string) administratorTable {
	var (
		IDColumn       = sqlite.IntegerColumn("id")
		UsernameColumn = sqlite.StringColumn("username")
		PasswordColumn = sqlite.StringColumn("password")
		allColumns     = sqlite.ColumnList{IDColumn, UsernameColumn, PasswordColumn}
		mutableColumns = sqlite.ColumnList{UsernameColumn, PasswordColumn}
	)

	return administratorTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Username: UsernameColumn,
		Password: PasswordColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
