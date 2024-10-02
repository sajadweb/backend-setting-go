package main

import (
	"go-category-app/config"
	"go-category-app/controllers"
	"go-category-app/repositories"
	"go-category-app/routes"
	"go-category-app/services"
	"log"
	"net/http"
)

func main() {
	config.ConnectDatabase()

	// Initialize repository and controller
	repo := repositories.NewCategoryRepository(config.DB)
	controller := controllers.NewCategoryController(repo)
	router := routes.SetupRoutes(controller)

	// Start HTTP server
	// go func() {
	// 	log.Println("Starting HTTP server on :8080")
	// 	log.Fatal(http.ListenAndServe(":8080", router))
	// }()

	// Start TCP server
	tcpService := services.NewTCPService(repo)
	tcpService.Start(":9090")  // Listens on port 9090 for TCP requests
}
