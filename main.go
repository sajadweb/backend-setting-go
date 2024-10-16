package main

import (
	"bakend-settings/common"
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
	fmt.Println("Tcp is running on", config.GetEnv("TCP_SERVER"))
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
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Read error:", err)
				break // Exit loop on error
			}
			// Log raw received data
			fmt.Println("Raw data received:", string(buf[:n]))

			// Process the received message
			request, convErr := common.MakeTcpRequest(conn,string(buf[:n]))
			if convErr != nil {
				fmt.Println("Error converting message:", convErr)
				break // Exit the read loop or handle as needed
			}
			// Check if the message is nil
			if request == nil {
				fmt.Println("Received nil message")
				break // Handle nil message case
			}
			go categoryController.Handle(request)
		}

	}
}
