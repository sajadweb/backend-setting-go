```go
package controllers

import (
	// "bakend-settings/models"
	"bakend-settings/common"
	"bakend-settings/models"
	"bakend-settings/services"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"regexp"

	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}
func (c *CategoryController) Handle(conn net.Conn) {
	// defer conn.Close()
	fmt.Println("TCP Server received connection")
	// defer conn.Close()
	// Buffer to read data from the connection
	buf := make([]byte, 2048)

	// Read the incoming data
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	// Slice the buffer to actual data size
	msg := buf[:n]
	fmt.Printf("Received message: %v\n", string(msg))
	re := regexp.MustCompile(`\{.*\}`)
	cleanedMsg := re.Find(msg)
	if cleanedMsg == nil {
		fmt.Println("Error: No valid JSON found in message")
		fmt.Fprintln(conn, "Invalid input")
		return
	}
	fmt.Printf("Received cleanedMsg: %v \n", string(cleanedMsg))

	// Define the expected structure for the incoming JSON
	var input struct {
		Action   string   `json:"pattern"`
		id   string   `json:"id"`
		Category  models.Category `json:"data"`
	}
	
	// Decode the incoming JSON
	err = json.Unmarshal(cleanedMsg, &input)
	if err != nil {
		fmt.Println("Error decoding input:", err)
		fmt.Fprintln(conn, "Invalid input")
		return
	}

	// Log the received action and data
	fmt.Printf("Received action: %s\n", input.Action)
	fmt.Printf("Received category: %+v\n", input.Category)

	// Context for MongoDB operations
	ctx := context.Background()

	// Handle the incoming action
	switch input.Action { 
	case "create.category":
		fmt.Println("Handling create.category")
		// result, err := c.service.Create(ctx, input.Category)
		// if err != nil {
		// 	common.SendResponse(conn, 500, "Error creating category", nil, 0, true)
		// 	return
		// }
		common.SendResponse(conn, 200, "Category retrieved", input.Category, 0, false)
	case "get.category":
		fmt.Println("Handling get.category")
		category, err := c.service.GetById(ctx, input.Category.ID)
		if err != nil { 
			common.SendResponse(conn, 500, "Error fetching category", nil, 0, true)
			return
		}
		json.NewEncoder(conn).Encode(category)
	case "update.category":
		fmt.Println("Handling update.category")
		_, err := c.service.Update(ctx, input.Category.ID, input.Category)
		if err != nil { 
			common.SendResponse(conn, 500, "Error updating category", nil, 0, true)
			return
		}
		fmt.Fprintln(conn, "Category updated successfully")
	case "delete.category":
		fmt.Println("Handling delete.category")
		_, err := c.service.Delete(ctx, input.Category.ID)
		if err != nil { 
			common.SendResponse(conn, 500, "Error deleting category", nil, 0, true)
			return
		}
		fmt.Fprintln(conn, "Category deleted successfully")
	case "find.all.category":
		fmt.Println("Handling find.all.category")
		categories, err := c.service.List(ctx)
		if err != nil { 
			common.SendResponse(conn, 500, "Error listing categories", nil, 0, true)
			return
		} 
		common.SendResponse(conn, 200, "Category retrieved", categories, 0, false)
	default:
		fmt.Fprintln(conn, "Unknown action")
		common.SendResponse(conn, 500, "Unknown action", nil, 0, true)
	}
}
