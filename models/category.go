package models

import "gopkg.in/mgo.v2/bson"

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
type Category struct {
	ID          bson.ObjectId    `bson:"_id,omitempty" json:"id"` //uuid.UUID `json:"id" db:"id"`
	//CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	//UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	Alias       string    `bson:"alias" json:"alias"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	Logo        string    `bson:"logo" json:"logo"`
	//ParentID    uuid.UUID `json:"parent_id" bson:"parent_id"`
}
//
//// String is not required by pop and may be deleted
//func (c Category) String() string {
//	jc, _ := json.Marshal(c)
//	return string(jc)
//}
//
// Categories is not required by pop and may be deleted
type Categories []Category
//
//// String is not required by pop and may be deleted
//func (c Categories) String() string {
//	jc, _ := json.Marshal(c)
//	return string(jc)
//}
//
//// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
//// This method is not required and may be deleted.
//func (c *Category) Validate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.Validate(
//		&validators.StringIsPresent{Field: c.Alias, Name: "Alias"},
//		&validators.StringIsPresent{Field: c.Title, Name: "Title"},
//		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
//		&validators.StringIsPresent{Field: c.Logo, Name: "Logo"},
//	), nil
//}
//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (c *Category) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (c *Category) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}
