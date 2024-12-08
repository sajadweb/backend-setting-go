package tests

import (
	"context"
	"fmt"
	"log"
	"testing"

	"bakend-settings/category/models"
	"bakend-settings/category/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCategoryCRUD(t *testing.T) {
	// Initialize real MongoDB client for integration test
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Get the collection and service
	collection := client.Database("testdb").Collection("categories")
	service := services.NewCategoryService(collection)

	// Test case for creating a category
	t.Run("Create", func(t *testing.T) {
		category := models.Category{
			ID:   primitive.NewObjectID(),
			Name: "Test Category",
		}
		_, err := service.Create(context.Background(), category)
		if err != nil {
			t.Fatalf("Error creating category: %v", err)
		}
	})
	t.Run("Get", func(t *testing.T) {
		category := models.Category{
			ID:   primitive.NewObjectID(),
			Name: "Test Category Get",
		}
		result, err := service.Create(context.Background(), category)
		fmt.Println("result is",result)
		if err != nil {
			t.Fatalf("Error creating category: %v", err)
		}
		// getCategory, err := service.GetById(context.Background(), result.InsertedID)
		// if err != nil {
		// 	t.Fatalf("Error getting category: %v", err)
		// }

		// if getCategory.Name != category.Name {
		// 	t.Errorf("Expected category name: %s, got: %s", category.Name, getCategory.Name)
		// }
	})
	// Additional tests for other CRUD operations can be placed here
}


// import (
// 	"testing"
// 	"context"

// 	"bakend-settings/services"
// 	"bakend-settings/models"

// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// func TestCategoryCRUD(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	collection := mt.Client.Database("testdb").Collection("categories")
// 	service := services.NewCategoryService(collection)

// 	t.Run("Create", func(t *testing.T) {
// 		_, err := service.Create(context.Background(), models.Category{Name: "Test Category"})
// 		if err != nil {
// 			t.Errorf("Error creating category: %v", err)
// 		}
// 	})

// 	// t.Run("Get", func(t *testing.T) {
// 	// 	category := models.Category{Name: "Test Category"}
// 	// 	result, err := service.Create(context.Background(), category)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error creating category: %v", err)
// 	// 	}

// 	// 	getCategory, err := service.GetById(context.Background(), result.InsertedID)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error getting category: %v", err)
// 	// 	}

// 	// 	if getCategory.Name != category.Name {
// 	// 		t.Errorf("Expected category name: %s, got: %s", category.Name, getCategory.Name)
// 	// 	}
// 	// })

// 	// t.Run("Update", func(t *testing.T) {
// 	// 	category := models.Category{Name: "Test Category"}
// 	// 	result, err := service.Create(context.Background(), category)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error creating category: %v", err)
// 	// 	}

// 	// 	category.Name = "Updated Category"
// 	// 	err = service.Update(context.Background(), result.InsertedID, category)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error updating category: %v", err)
// 	// 	}

// 	// 	getCategory, err := service.GetById(context.Background(), result.InsertedID)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error getting category: %v", err)
// 	// 	}

// 	// 	if getCategory.Name != category.Name {
// 	// 		t.Errorf("Expected updated category name: %s, got: %s", category.Name, getCategory.Name)
// 	// 	}
// 	// })

// 	// t.Run("Delete", func(t *testing.T) {
// 	// 	category := models.Category{Name: "Test Category"}
// 	// 	result, err := service.Create(context.Background(), category)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error creating category: %v", err)
// 	// 	}

// 	// 	err = service.Delete(context.Background(), result.InsertedID)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error deleting category: %v", err)
// 	// 	}

// 	// 	_, err = service.GetById(context.Background(), result.InsertedID)
// 	// 	if err == nil {
// 	// 		t.Error("Expected category to be deleted")
// 	// 	}
// 	// })

// 	// t.Run("List", func(t *testing.T) {
// 	// 	categories, err := service.List(context.Background())
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error listing categories: %v", err)
// 	// 	}

// 	// 	if len(categories) == 0 {
// 	// 		t.Error("Expected at least one category in the list")
// 	// 	}
// 	// })

// 	defer mt.Close()
// }
