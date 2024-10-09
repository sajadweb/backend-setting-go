package main

import (
	"bakend-settings/config"
	"bakend-settings/controllers"
	"bakend-settings/services"
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.LoadEnv()
	fmt.Println("MONGODB_URI is", config.GetEnv("MONGODB_URI"))
	clientOptions := options.Client().ApplyURI(config.GetEnv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(config.GetEnv("MONGODB_DB")).Collection("categories")
	categoryService := services.NewCategoryService(collection)
	categoryController := controllers.NewCategoryController(categoryService)

	listener, err := net.Listen("tcp", config.GetEnv("TCP_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Read error:", err)
			log.Fatal(err)
		} 
	 
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}
		go categoryController.Handle(conn)
		conn.Close()
	}
}
