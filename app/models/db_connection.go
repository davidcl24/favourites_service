// This package represents the entity in the database and the CRUD operations.
package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// A wrapper for the SQL connection to handle it in an easier way.
type DB struct {
	Conn *sql.DB
}

// Creates a new instance for the DB struct
func NewDB(conn *sql.DB) *DB {
	return &DB{Conn: conn}
}
