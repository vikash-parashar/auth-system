// main.go
package main

import (
	"auth-system/config"
	"auth-system/routes"
)

func main() {
	// Initialize database connection
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	// Set up the API routes
	r := routes.SetupRouter()

	// Start the server
	r.Run(":8080")
}
