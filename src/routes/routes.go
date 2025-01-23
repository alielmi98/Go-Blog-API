package routes

import (
	"github.com/alielmi98/Go-Blog-API/controllers"
	"github.com/alielmi98/Go-Blog-API/repository"
	"github.com/alielmi98/Go-Blog-API/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {

	// Create the repositories
	postRepository := repository.NewPostRepository(db)

	// Create the services
	postService := services.NewPostService(postRepository)

	// Create the controllers
	postController := controllers.NewPostController(postService)

	// Define the routes
	api := r.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", postController.CreatePost)
			posts.GET("/", postController.GetAllPosts)
			posts.PUT("/:id", postController.UpdatePost)
			posts.DELETE("/:id", postController.DeletePost)
			posts.GET("/:id", postController.GetPostById)
			posts.GET("/filter", postController.GetByFilter)

		}
	}
}
