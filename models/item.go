package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

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
type Item struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	//UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	Alias       string    `json:"alias" bson:"alias"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Pictures    string    `json:"pictures" bson:"pictures"`
	Price       int       `json:"price" bson:"price"`
	Count       int       `json:"count" bson:"count"`
	//CategoryID  uuid.UUID `json:"category_id" bson:"category_id"`
}
//
//// String is not required by pop and may be deleted
//func (i Item) String() string {
//	ji, _ := json.Marshal(i)
//	return string(ji)
//}
//
// Items is not required by pop and may be deleted
type Items []Item
//
//// String is not required by pop and may be deleted
//func (i Items) String() string {
//	ji, _ := json.Marshal(i)
//	return string(ji)
//}
//
//// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
//// This method is not required and may be deleted.
//func (i *Item) Validate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.Validate(
//		&validators.StringIsPresent{Field: i.Alias, Name: "Alias"},
//		&validators.StringIsPresent{Field: i.Title, Name: "Title"},
//		&validators.StringIsPresent{Field: i.Description, Name: "Description"},
//		&validators.StringIsPresent{Field: i.Pictures, Name: "Pictures"},
//		&validators.IntIsPresent{Field: i.Price, Name: "Price"},
//		&validators.IntIsPresent{Field: i.Count, Name: "Count"},
//	), nil
//}
//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (i *Item) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (i *Item) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
