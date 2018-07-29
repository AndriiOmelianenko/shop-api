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

// ItemsList default implementation.
func ItemsList(w http.ResponseWriter, r *http.Request) {
	items := models.Items{}
	err := dao.DB.C(dao.COLLECTION_ITEMS).Find(bson.M{}).All(&items)
	if err != nil {
		fmt.Println("error getting list of items:", err)
		fmt.Fprintln(w, "error getting list of items:", err)
		return
	}
	json.NewEncoder(w).Encode(items)
}

// ItemsIndex default implementation.
func ItemsIndex(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	params := mux.Vars(r)
	err := dao.DB.C(dao.COLLECTION_ITEMS).FindId(bson.ObjectIdHex(params["item"])).One(&item)
	if err != nil {
		fmt.Println("error getting specific item:", err)
		fmt.Fprintln(w, "error getting specific item:", err)
		return
	}
	json.NewEncoder(w).Encode(item)
}
