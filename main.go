// main.go
package main

import (
	"auth-system/config"
	"auth-system/routes"
	"database/sql"
)

var (
	DB *sql.DB
)

func main() {
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	DB = db
	defer db.Close()

	// Set up the API routes
	r := routes.SetupRouter()

	// Start the server
	r.Run(":8080")
}
