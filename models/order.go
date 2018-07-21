package models

//
//import (
//	"encoding/json"
//	"time"
//
//	"github.com/gobuffalo/pop"
//	"github.com/gobuffalo/validate"
//	"github.com/gobuffalo/validate/validators"
//)
//
type Order struct {
	ID int `json:"id" bson:"_id"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
	//UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Status string `json:"status" bson:"status"`
	Sum    int    `json:"sum" bson:"sum"`
}

//
//// String is not required by pop and may be deleted
//func (o Order) String() string {
//	jo, _ := json.Marshal(o)
//	return string(jo)
//}
//
//// Orders is not required by pop and may be deleted
type Orders []Order

//
//// String is not required by pop and may be deleted
//func (o Orders) String() string {
//	jo, _ := json.Marshal(o)
//	return string(jo)
//}
//
//// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
//// This method is not required and may be deleted.
//func (o *Order) Validate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.Validate(
//		&validators.IntIsPresent{Field: o.ID, Name: "ID"},
//		&validators.StringIsPresent{Field: o.Status, Name: "Status"},
//		&validators.IntIsPresent{Field: o.Sum, Name: "Sum"},
//	), nil
//}
//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (o *Order) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (o *Order) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
