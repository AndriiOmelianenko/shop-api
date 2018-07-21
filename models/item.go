package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	Alias       string        `json:"alias" bson:"alias"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Pictures    string        `json:"pictures" bson:"pictures"`
	Price       int           `json:"price" bson:"price"`
	Count       int           `json:"count" bson:"count"`
	CategoryID  bson.ObjectId `json:"category_id" bson:"category_id"`
}

type Items []Item
