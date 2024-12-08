package services

import (
	"context"
	"log"

	"bakend-settings/category/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService struct {
	collection *mongo.Collection
}

func NewCategoryService(collection *mongo.Collection) *CategoryService {
	return &CategoryService{
		collection: collection,
	}
}

func (s *CategoryService) Create(ctx context.Context, category models.Category) (*mongo.InsertOneResult, error) {
	category.ID = primitive.NewObjectID()
	return s.collection.InsertOne(ctx, category)
}

func (s *CategoryService) GetById(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	var category models.Category
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	return &category, err
}

func (s *CategoryService) Update(ctx context.Context, id primitive.ObjectID, category models.Category) (*mongo.UpdateResult, error) {
	return s.collection.UpdateByID(ctx, id, bson.M{"$set": category})
}

func (s *CategoryService) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(ctx, bson.M{"_id": id})
}

func (s *CategoryService) List(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error finding categories:", err)
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
