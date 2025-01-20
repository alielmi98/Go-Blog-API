package main

import (
	"github.com/alielmi98/Go-Blog-API/config"
	"github.com/alielmi98/Go-Blog-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin engine
	r := gin.Default()

	// Initialize the database connection
	db, _ := config.ConnectDB()

	// Register the routes
	routes.RegisterRoutes(r, db)

	// Start the server
	r.Run(":8080")
}
