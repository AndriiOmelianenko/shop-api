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
// curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/orders -d '{"status": "new"}'
func OrdersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	order := models.Order{}
	err := decoder.Decode(&order)
	if err != nil {
		fmt.Println("error decoding json:", err)
		fmt.Fprintln(w, "error decoding json:", err)
		return
	}
	currentTime := time.Now()
	order.ID = bson.NewObjectId()
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
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/orders
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
// curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/orders/<orderID>
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

//
//// curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/orders/3/item -d '{"item_id": "fa0f13d1-af55-4823-8f85-7e3284316b70", "item_cnt": 5}'
//// OrdersUpdate default implementation.
func OrdersCreateItem(w http.ResponseWriter, r *http.Request) {
	//	decoder := json.NewDecoder(c.Request().Body)
	//	ordered := models.Ordered{}
	//	err := decoder.Decode(&ordered)
	//	if err != nil {
	//		return c.Render(404, r.String("ERROR: decoding json: %v", err))
	//	}
	//	orderID, err := strconv.Atoi(c.Param("order"))
	//	if err != nil {
	//		return c.Render(404, r.String("ERROR: converting orderID to int: %v", err))
	//	}
	//	ordered.OrderID = orderID
	//
	//	// get item price and calculate ordered price based on item_price * item_count
	//	item := models.Item{}
	//	err = models.DB.Find(&item, ordered.ItemID)
	//	if err != nil {
	//		return c.Render(404, r.String("ERROR: cant find item in database: %v", err))
	//	}
	//	ordered.ItemSum = ordered.ItemCnt * item.Price
	//
	//	err = models.DB.Create(&ordered)
	//	if err != nil {
	//		return c.Render(404, r.String("ERROR: creating ordered record: %v", err))
	//	}
	//	return c.Render(200, r.String("Item added to order: %v", ordered))
	fmt.Fprintln(w, "not implemented yet !")
}
