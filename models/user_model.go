package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" swaggerignore:"true"`
	Name        string             `json:"name,omitempty" validate:"required"`
	DOB         string             `json:"dob,omitempty" validate:"required"`
	Address     string             `json:"address,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt" swaggerignore:"true"`
}
