package models

//
//import (
//	"encoding/json"
//	"time"
//
//	"github.com/gobuffalo/pop"
//	"github.com/gobuffalo/uuid"
//	"github.com/gobuffalo/validate"
//	"github.com/gobuffalo/validate/validators"
//)
//
type Ordered struct {
	ID int `json:"id" bson:"id"`
	//CreatedAt time.Time `json:"created_at" bson:"created_at"`
	//UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	OrderID int `json:"order_id" bson:"order_id"`
	//ItemID    uuid.UUID `json:"item_id" bson:"item_id"`
	ItemCnt int `json:"item_cnt" bson:"item_cnt"`
	ItemSum int `json:"item_sum" bson:"item_sum"`
}

//
//// String is not required by pop and may be deleted
//func (o Ordered) String() string {
//	jo, _ := json.Marshal(o)
//	return string(jo)
//}
//
//// Ordereds is not required by pop and may be deleted
type Ordereds []Ordered

//
//// String is not required by pop and may be deleted
//func (o Ordereds) String() string {
//	jo, _ := json.Marshal(o)
//	return string(jo)
//}
//
//// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
//// This method is not required and may be deleted.
//func (o *Ordered) Validate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.Validate(
//		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
//		&validators.IntIsPresent{Field: o.OrderID, Name: "OrderID"},
//		&validators.IntIsPresent{Field: o.ItemCnt, Name: "ItemCnt"},
//		&validators.IntIsPresent{Field: o.ItemSum, Name: "ItemSum"},
//	), nil
//}
//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (o *Ordered) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (o *Ordered) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
