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

var UsersDiscussions = newUsersDiscussionsTable("", "users_discussions", "")

type usersDiscussionsTable struct {
	sqlite.Table

	// Columns
	UserID       sqlite.ColumnInteger
	DiscussionID sqlite.ColumnInteger

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type UsersDiscussionsTable struct {
	usersDiscussionsTable

	EXCLUDED usersDiscussionsTable
}

// AS creates new UsersDiscussionsTable with assigned alias
func (a UsersDiscussionsTable) AS(alias string) *UsersDiscussionsTable {
	return newUsersDiscussionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersDiscussionsTable with assigned schema name
func (a UsersDiscussionsTable) FromSchema(schemaName string) *UsersDiscussionsTable {
	return newUsersDiscussionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersDiscussionsTable with assigned table prefix
func (a UsersDiscussionsTable) WithPrefix(prefix string) *UsersDiscussionsTable {
	return newUsersDiscussionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersDiscussionsTable with assigned table suffix
func (a UsersDiscussionsTable) WithSuffix(suffix string) *UsersDiscussionsTable {
	return newUsersDiscussionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersDiscussionsTable(schemaName, tableName, alias string) *UsersDiscussionsTable {
	return &UsersDiscussionsTable{
		usersDiscussionsTable: newUsersDiscussionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newUsersDiscussionsTableImpl("", "excluded", ""),
	}
}

func newUsersDiscussionsTableImpl(schemaName, tableName, alias string) usersDiscussionsTable {
	var (
		UserIDColumn       = sqlite.IntegerColumn("user_id")
		DiscussionIDColumn = sqlite.IntegerColumn("discussion_id")
		allColumns         = sqlite.ColumnList{UserIDColumn, DiscussionIDColumn}
		mutableColumns     = sqlite.ColumnList{UserIDColumn, DiscussionIDColumn}
	)

	return usersDiscussionsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:       UserIDColumn,
		DiscussionID: DiscussionIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
