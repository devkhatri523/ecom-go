package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name" validate:"required,email"`
	LastName  string             `bson:"last_name" validate:"required,email"`
	Email     string             `bson:"email" validate:"required,email"`
	Address   Address            `bson:"address" validate:"required,email"`
}
