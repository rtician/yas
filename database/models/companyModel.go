package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"yas/types"
)

type Company struct {
	Id          primitive.ObjectID `bson:"id" json:"_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	Active      bool               `bson:"active" json:"active"`
	Name        string             `bson:"name" json:"name"`
	Document    string             `bson:"document" json:"document"`
	Email       string             `bson:"email" json:"email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Address     *types.Address     `bson:"address" json:"address"`
}

func NewCompany() *Company {
	return &Company{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Address:   &types.Address{},
	}
}
