// config/config.go
package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	DBDriver     = "postgres"
	DBDataSource = "postgresql://postgres:admin@localhost/auth_system_db?sslmode=disable"
)

// InitDB initializes the database connection and returns a pointer to the database.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open(DBDriver, DBDataSource)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
