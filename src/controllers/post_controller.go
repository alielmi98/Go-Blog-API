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

func (pc *PostController) UpdatePost(c *gin.Context) {
	var updatePostDTO dto.UpdatePostDTO
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&updatePostDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	post, err := pc.postService.UpdatePost(c, id, &updatePostDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	err = pc.postService.DeletePost(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted post successfully")
}

func (pc *PostController) GetPostById(c *gin.Context) {
	id := c.Params.ByName("id")

	post, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postService.GetAllPosts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)

}
