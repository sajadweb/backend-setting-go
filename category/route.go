package category

import (
	"bakend-settings/category/controllers"
	"bakend-settings/category/services"
	"bakend-settings/config"
	"bakend-settings/microservices/tcp"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CateogyMain(app *tcp.TcpServer) {
    // config.LoadEnv()
    clientOptions := options.Client().ApplyURI(config.GetEnv("MONGODB_URI"))
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }
    
    collection := client.Database(config.GetEnv("MONGODB_DB")).Collection("categories")
    categoryService := services.NewCategoryService(collection)
    categoryController := controllers.NewCategoryController(categoryService)
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
    app.Pattern("create.category", categoryController.Create)
    app.Pattern("find.all.categories", categoryController.FindAll)

	fmt.Println("Service started and MongoDB connection is live")
	 // Ensure proper disconnection only when needed, not too early
	//defer client.Disconnect(context.Background())
}
