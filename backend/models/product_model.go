package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reviews struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Rating    int                `json:"rating" bson:"rating" validate:"required"`
	Comment   string             `json:"comment" bson:"comment" validate:"required"`
	User      primitive.ObjectID `json:"user" bson:"user_id" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Image       string             `json:"image" bson:"image" validate:"required"`
	Brand       string             `json:"brand" bson:"brand" validate:"required"`
	Quantity    int                `json:"quantity" bson:"quantity" validate:"required"`
	Category    primitive.ObjectID `json:"category" bson:"category_id" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Reviews     Reviews            `bson:"inline"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
