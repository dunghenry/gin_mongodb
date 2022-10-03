package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title"`
	Des       string             `json:"des"`
	Completed bool               `json:"completed"`
}
