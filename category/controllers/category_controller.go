package controllers

import (
	"bakend-settings/category/models"
	"bakend-settings/category/services"
	"bakend-settings/microservices/tcp/common"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}
func (c *CategoryController) FindAll(req *common.TcpRequest) {
	// Context for MongoDB operations
	ctx := context.Background()
	fmt.Println(ctx)
	
	categories, err := c.service.List(ctx)
	if err != nil {
		fmt.Println(err)
		common.SendResponse(req, 500, nil, "Error listing categories1", true)
		return
	}
	common.SendResponse(req, 200, categories, "Category retrieved2", false)
}
func (c *CategoryController) Create(req *common.TcpRequest) {
	// Context for MongoDB operations
	ctx := context.Background()
	fmt.Println(ctx)

	category := models.Category{
		Name: req.Data["name"].(string),
		Icon: req.Data["icon"].(string),
	}

	result, err := c.service.Create(ctx, category)
	if err != nil {
		common.SendResponse(req, 500, nil, "Error creating category", true)
		return
	}
	category.ID = result.InsertedID.(primitive.ObjectID)
	common.SendResponse(req, 200, struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}{ID: category.ID.Hex(), Name: category.Name, Icon: category.Icon}, "Category created successfully", false)
}

