package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name" validate:"required,max=32"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required,max=32"`
}