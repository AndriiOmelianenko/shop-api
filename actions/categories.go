package actions

import (
	"net/http"
	"encoding/json"
	"github.com/AndriiOmelianenko/shop-api/models"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/AndriiOmelianenko/shop-api/dao"
	"github.com/gorilla/mux"
)

// CategoriesList default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories
func CategoriesList(w http.ResponseWriter, r *http.Request) {
	categories := models.Categories{}
	err := dao.DB.C(dao.COLLECTION_CATEGORIES).Find(bson.M{}).All(&categories)
	if err != nil {
		fmt.Println("error getting list of categories:", err)
	}
	json.NewEncoder(w).Encode(categories)
}

// CategoriesIndex default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories/<categoryID>
func CategoriesIndex(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	params := mux.Vars(r)
	err := dao.DB.C(dao.COLLECTION_CATEGORIES).FindId(bson.ObjectIdHex(params["category"])).One(&category)
	if err != nil {
		fmt.Println("error getting specific category:", err)
	}
	json.NewEncoder(w).Encode(category)
}
