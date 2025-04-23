package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty" binding:"required"`
	Email     string             `json:"email" bson:"email,omitempty" binding:"required,email"`
	Birthdate time.Time          `json:"birthdate" bson:"birthdate,omitempty" binding:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
