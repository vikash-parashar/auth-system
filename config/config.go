// config/config.go
package config

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection and returns a pointer to the database.
func InitDB() (*gorm.DB, error) {
	// Define your PostgreSQL database connection string
	dsn := "user=postgres password=admin dbname=auth_system_db host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return db, err
	}
	return db, nil
}
