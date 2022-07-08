package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"create_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"update_at" json:"updated_at"`
	Name      string             `bson:"name" json:"name"`
	Password  string             `bson:"password" json:"password"`
	Email     string             `bson:"email" json:"email"`
}

func NewUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
