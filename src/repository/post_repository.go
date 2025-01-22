package repository

import (
	"context"

	"github.com/alielmi98/Go-Blog-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	UpdatePost(ctx context.Context, id string, post *models.Post) error
	DeletePost(ctx context.Context, id string) error
	GetPostByID(ctx context.Context, id string) (*models.Post, error)
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
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
	result, err := r.collection.InsertOne(ctx, post)
	post.ID = result.InsertedID.(primitive.ObjectID)
	return err
}
func (r *postRepository) UpdatePost(ctx context.Context, id string, post *models.Post) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": post})

	return err
}

func (r *postRepository) DeletePost(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})

	return err
}

func (r *postRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var post models.Post
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
