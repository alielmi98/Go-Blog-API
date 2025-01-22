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
	UpdatePost(ctx context.Context, id string, post *dto.UpdatePostDTO) (*models.Post, error)
	DeletePost(ctx context.Context, id string) error
	GetPostByID(ctx context.Context, id string) (*models.Post, error)
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
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

func (s *postService) UpdatePost(ctx context.Context, id string, post *dto.UpdatePostDTO) (*models.Post, error) {
	existingPost, _ := s.postRepository.GetPostByID(ctx, id)
	udatePost := &models.Post{
		ID:        existingPost.ID,
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		Tags:      post.Tags,
		CreatedAt: existingPost.CreatedAt,
		UpdatedAt: time.Now(),
	}
	err := s.postRepository.UpdatePost(ctx, id, udatePost)
	return udatePost, err
}

func (s *postService) DeletePost(ctx context.Context, id string) error {
	err := s.postRepository.DeletePost(ctx, id)
	return err
}

func (s *postService) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	post, err := s.postRepository.GetPostByID(ctx, id)
	return post, err

}
func (s postService) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	return s.postRepository.GetAllPosts(ctx)
}
