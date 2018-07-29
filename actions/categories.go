package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AndriiOmelianenko/shop-api/dao"
	"github.com/AndriiOmelianenko/shop-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// CategoriesList default implementation.
func CategoriesList(w http.ResponseWriter, r *http.Request) {
	categories := models.Categories{}
	err := dao.DB.C(dao.COLLECTION_CATEGORIES).Find(bson.M{}).All(&categories)
	if err != nil {
		fmt.Println("error getting list of categories:", err)
		fmt.Fprintln(w, "error getting list of categories:", err)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// CategoriesIndex default implementation.
func CategoriesIndex(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	params := mux.Vars(r)
	err := dao.DB.C(dao.COLLECTION_CATEGORIES).FindId(bson.ObjectIdHex(params["category"])).One(&category)
	if err != nil {
		fmt.Println("error getting specific category:", err)
		fmt.Fprintln(w, "error getting specific category:", err)
		return
	}
	json.NewEncoder(w).Encode(category)
}
