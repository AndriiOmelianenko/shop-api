package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
	Alias       string        `bson:"alias" json:"alias"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Logo        string        `bson:"logo" json:"logo"`
	ParentID    bson.ObjectId `json:"parent_id" bson:"parent_id,omitempty"`
}

type Categories []Category
