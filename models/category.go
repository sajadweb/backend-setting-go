package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Icon     string             `bson:"icon" json:"icon"`
	ParentID *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	Children []*Category        `bson:"children,omitempty" json:"children,omitempty"`
}
