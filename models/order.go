package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	Status    string        `json:"status" bson:"status"`
	Sum       int           `json:"sum" bson:"sum"`
}

type Orders []Order
