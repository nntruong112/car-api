package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Make  string             `bson:"make" json:"make"`
	Model string             `bson:"model" json:"model"`
	Year  int                `bson:"year" json:"year"`
}
