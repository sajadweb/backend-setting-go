package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`         // MongoDB ObjectID is mapped to "id"
    Name     string             `bson:"name" json:"name"`                // Name field is mapped to "name"
    Icon     string             `bson:"icon" json:"icon"`                // Icon field is mapped to "icon"
    ParentID *primitive.ObjectID `bson:"parentId,omitempty" json:"parentId"` // Optional ParentID is mapped to "parentId"
    Children []*Category        `bson:"children,omitempty" json:"children"` // Optional Children field is mapped to "children"

}
