package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/AndriiOmelianenko/shop-api/dao"
	"github.com/AndriiOmelianenko/shop-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// OrdersCreate default implementation.
func OrdersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	order := models.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		fmt.Println("error decoding json:", err)
		fmt.Fprintln(w, "error decoding json:", err)
		return
	}
	order.ID = bson.NewObjectId()
	currentTime := time.Now()
	order.CreatedAt = currentTime
	order.UpdatedAt = currentTime

	err = dao.DB.C(dao.COLLECTION_ORDERS).Insert(&order)
	if err != nil {
		fmt.Println("error adding order to mongo:", err)
		fmt.Fprintln(w, "error adding order to mongo:", err)
		return
	}

	fmt.Fprint(w, "Order created: ")
	json.NewEncoder(w).Encode(order)
}

// OrdersList default implementation.
func OrdersList(w http.ResponseWriter, r *http.Request) {
	orders := models.Orders{}
	err := dao.DB.C(dao.COLLECTION_ORDERS).Find(bson.M{}).All(&orders)
	if err != nil {
		fmt.Println("error getting list of orders:", err)
		fmt.Fprintln(w, "error getting list of orders:", err)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

// OrdersIndex default implementation.
func OrdersIndex(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}
	params := mux.Vars(r)
	err := dao.DB.C(dao.COLLECTION_ORDERS).FindId(bson.ObjectIdHex(params["order"])).One(&order)
	if err != nil {
		fmt.Println("error getting specific order:", err)
		fmt.Fprintln(w, "error getting specific order:", err)
		return
	}
	json.NewEncoder(w).Encode(order)
}

// OrdersUpdate default implementation.
func OrdersCreateItem(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	ordered := models.Ordered{}
	err := decoder.Decode(&ordered)
	if err != nil {
		fmt.Println("error decoding json:", err)
		fmt.Fprintln(w, "error decoding json:", err)
		return
	}

	params := mux.Vars(r)

	orderID := bson.ObjectIdHex(params["order"])
	if err != nil {
		fmt.Println("error converting orderID:", err)
		fmt.Fprintln(w, "error converting orderID:", err)
		return
	}
	ordered.OrderID = orderID

	// get order information
	order := models.Order{}
	err = dao.DB.C(dao.COLLECTION_ORDERS).FindId(bson.ObjectId(ordered.OrderID)).One(&order)
	if err != nil {
		fmt.Println("cant find order in database:", err)
		fmt.Fprintln(w, "cant find order in database:", err)
		return
	}

	// get item information
	item := models.Item{}
	err = dao.DB.C(dao.COLLECTION_ITEMS).FindId(bson.ObjectId(ordered.ItemID)).One(&item)
	if err != nil {
		fmt.Println("cant find item in database:", err)
		fmt.Fprintln(w, "cant find item in database:", err)
		return
	}

	ordered.ID = bson.NewObjectId()
	currentTime := time.Now()
	ordered.CreatedAt = currentTime
	ordered.UpdatedAt = currentTime

	// ordered price based on item_price * item_count
	ordered.ItemSum = ordered.ItemCnt * item.Price

	// update total sum in order
	order.Sum = order.Sum + ordered.ItemSum

	// insert new ordered
	err = dao.DB.C(dao.COLLECTION_ORDEREDS).Insert(&ordered)
	if err != nil {
		fmt.Println("error adding ordered to mongo:", err)
		fmt.Fprintln(w, "error adding ordered to mongo:", err)
		return
	}

	// update order
	err = dao.DB.C(dao.COLLECTION_ORDERS).UpdateId(bson.ObjectId(order.ID), &order)
	if err != nil {
		fmt.Println("error updating order to mongo:", err)
		fmt.Fprintln(w, "error updating order to mongo:", err)
		return
	}

	fmt.Fprint(w, "Item added to order: ")
	json.NewEncoder(w).Encode(ordered)

	fmt.Fprint(w, "Order updated: ")
	json.NewEncoder(w).Encode(order)
}
