package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Ordered struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	OrderID   int           `json:"order_id" bson:"order_id"`
	ItemID    bson.ObjectId `json:"item_id" bson:"item_id"`
	ItemCnt   int           `json:"item_cnt" bson:"item_cnt"`
	ItemSum   int           `json:"item_sum" bson:"item_sum"`
}

type Ordereds []Ordered
