package controllers

import (
	"bakend-settings/common"
	"bakend-settings/models"

	// "encoding/json"

	// "bakend-settings/models"
	"bakend-settings/services"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "fmt"
	// "net"
	// "regexp"
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
func (c *CategoryController) Handle(req *common.TcpRequest) {

	// err = json.Unmarshal(req.Data, &category)
	// if err != nil {
	// 	common.SendResponse(conn, 500, "Error unmarshalling JSON", nil, 0, true)
	// 	return
	// }
	// Context for MongoDB operations
	ctx := context.Background()

	// Handle the incoming action
	switch req.Pattern {

	case "create.category":
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
	case "find.all.categories":
		categories, err := c.service.List(ctx)
		if err != nil {
			common.SendResponse(req, 500, nil, "Error listing categories", true)
			return
		}
		common.SendResponse(req, 200, categories, "Category retrieved", false)
	}
}
