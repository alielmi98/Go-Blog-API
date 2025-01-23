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

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new blog post with the provided details
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dto.CreatePostDTO true "Post data"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts [post]
func (pc *PostController) CreatePost(c *gin.Context) {
	var postDTO dto.CreatePostDTO
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	post, err := pc.postService.CreatePost(c, &postDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// UpdatePost godoc
// @Summary Update an existing post
// @Description Update a post by ID with the provided details
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body dto.UpdatePostDTO true "Updated Post data"
// @Success 200 {object} models.Post
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/{id} [put]
func (pc *PostController) UpdatePost(c *gin.Context) {
	var updatePostDTO dto.UpdatePostDTO
	id := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&updatePostDTO); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	_, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	post, err := pc.postService.UpdatePost(c, id, &updatePostDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags posts
// @Param id path string true "Post ID"
// @Success 200 {string} string "Deleted post successfully"
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /posts/{id} [delete]
func (pc *PostController) DeletePost(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "post not found"})
		return
	}

	err = pc.postService.DeletePost(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted post successfully")
}

// GetPostById godoc
// @Summary Get a post by ID
// @Description Retrieve a post by its ID
// @Tags posts
// @Param id path string true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]interface{}
// @Router /posts/{id} [get]
func (pc *PostController) GetPostById(c *gin.Context) {
	id := c.Params.ByName("id")

	post, err := pc.postService.GetPostByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetAllPosts godoc
// @Summary Get all posts
// @Description Retrieve a list of all posts
// @Tags posts
// @Success 200 {array} models.Post
// @Failure 500 {object} map[string]interface{}
// @Router /posts [get]
func (pc *PostController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postService.GetAllPosts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)

}
