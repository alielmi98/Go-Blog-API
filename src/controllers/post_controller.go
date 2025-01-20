package controllers

import (
	"net/http"

	"github.com/alielmi98/Go-Blog-API/dto"
	"github.com/alielmi98/Go-Blog-API/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var postDTO dto.CreatePostDTO
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := pc.postService.CreatePost(c, &postDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// Implement other CRUD operations similar to CreatePost
