package main

import (
	"log"

	"github.com/alielmi98/Go-Blog-API/config"
	"github.com/alielmi98/Go-Blog-API/docs"
	"github.com/alielmi98/Go-Blog-API/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Create a new Gin engine
	r := gin.Default()

	// Initialize the database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Register the routes
	routes.RegisterRoutes(r, db)
	RegisterSwagger(r)
	// Start the server
	r.Run(":8080")

}

func RegisterSwagger(r *gin.Engine) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
