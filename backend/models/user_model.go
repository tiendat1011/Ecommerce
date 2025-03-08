package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string        `json:"username" bson:"username" validate:"required"`
	Email     string        `json:"email" bson:"email" validate:"required,email"`
	Password  string        `json:"password" bson:"password" validate:"required,min=6"`
	IsAdmin   bool          `json:"is_admin" bson:"is_admin,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}
