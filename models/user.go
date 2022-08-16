package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Dob         string             `json:"dob" bson:"dob"`
	Address     Address            `json:"address" bson:"address"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

type Address struct {
	City     string `json:"city,omitempty"  bson:"city,omitempty"`
	Country  string `json:"country,omitempty" bson:"country,omitempty"`
	StreetNo string `json:"streetNo,omitempty" bson:"streetNo,omitempty"`
	Pincode  string `json:"pincode,omitempty" bson:"pincode,omitempty"`
}
