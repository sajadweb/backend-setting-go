package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Icon     string             `bson:"icon,omitempty"`
	ParentID *primitive.ObjectID `bson:"parentId,omitempty"`
	Children []Category         `bson:"children,omitempty"`
}
