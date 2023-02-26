package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        *string            `json:"name"`
	Author      *string            `json:"author"`
	Publication *string            `json:"pub"`
}