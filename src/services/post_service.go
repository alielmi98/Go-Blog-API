// service/post_service.go
package services

import (
	"context"
	"time"

	"github.com/alielmi98/Go-Blog-API/dto"
	"github.com/alielmi98/Go-Blog-API/models"

	"github.com/alielmi98/Go-Blog-API/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post *dto.CreatePostDTO) (*models.Post, error)
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (s *postService) CreatePost(ctx context.Context, post *dto.CreatePostDTO) (*models.Post, error) {
	newPost := &models.Post{
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		Tags:      post.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.postRepository.CreatePost(ctx, newPost)
	return newPost, err
}
