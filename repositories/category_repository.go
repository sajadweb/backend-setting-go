package repositories

import (
	"context"
	"go-category-tcp-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CategoryRepository struct {
	DB *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{DB: db.Collection("categories")}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category models.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id primitive.ObjectID) (models.Category, error) {
	var category models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.DB.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	return category, err
}

func (r *CategoryRepository) Create(category models.Category) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.DB.InsertOne(ctx, category)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *CategoryRepository) Update(id primitive.ObjectID, category models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.DB.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"name":      category.Name,
			"icon":      category.Icon,
			"parent_id": category.ParentID,
		},
	})
	return err
}

func (r *CategoryRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.DB.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
