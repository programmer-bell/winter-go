package repository

import (
	"context"
	"time"

	"note_api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) *BlogRepository {
	return &BlogRepository{
		collection: db.Collection("blogs"),
	}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(context.Background(), blog)
	return err
}

func (r *BlogRepository) GetAll() ([]models.Blog, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var blogs []models.Blog
	if err = cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (r *BlogRepository) GetByID(id string) (*models.Blog, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var blog models.Blog
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&blog)
	return &blog, err
}

func (r *BlogRepository) Update(id string, blog *models.Blog) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	blog.UpdatedAt = time.Now()

	_, err = r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": blog},
	)
	return err
}

func (r *BlogRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}
