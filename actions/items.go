package actions

import (
	"fmt"
	"net/http"
	"github.com/AndriiOmelianenko/shop-api/dao"
	"gopkg.in/mgo.v2/bson"
	"github.com/AndriiOmelianenko/shop-api/models"
	"encoding/json"
	"github.com/gorilla/mux"
)

// ItemsList default implementation.
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/items
func ItemsList(w http.ResponseWriter, r *http.Request) {
	items := models.Items{}
	err := dao.DB.C(dao.COLLECTION_ITEMS).Find(bson.M{}).All(&items)
	if err != nil {
		fmt.Println("error getting list of items:", err)
	}
	json.NewEncoder(w).Encode(items)
}

// ItemsIndex default implementation.
//curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/items/<itemID>
func ItemsIndex(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	params := mux.Vars(r)
	err := dao.DB.C(dao.COLLECTION_ITEMS).FindId(bson.ObjectIdHex(params["item"])).One(&item)
	if err != nil {
		fmt.Println("error getting specific item:", err)
	}
	json.NewEncoder(w).Encode(item)
}
