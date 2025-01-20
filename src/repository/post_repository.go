package repository

import (
	"context"

	"github.com/alielmi98/Go-Blog-API/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
}

type postRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) PostRepository {
	return &postRepository{
		collection: db.Collection("posts"),
	}
}

func (r *postRepository) CreatePost(ctx context.Context, post *models.Post) error {
	_, err := r.collection.InsertOne(ctx, post)
	return err
}
