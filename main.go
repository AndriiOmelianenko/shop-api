package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/AndriiOmelianenko/shop-api/actions"
	"github.com/AndriiOmelianenko/shop-api/dao"
	"fmt"
)

// main is the starting point to shop-api application.
//func main() {
//	app := actions.App()
//	if err := app.Serve(); err != nil {
//		log.Fatal(err)
//	}
//}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", actions.HomeHandler).Methods("GET")
	
	router.HandleFunc("/items", actions.ItemsList).Methods("GET")
	router.HandleFunc("/items/{item}", actions.ItemsIndex).Methods("GET")
	
	router.HandleFunc("/categories", actions.CategoriesList).Methods("GET")
	router.HandleFunc("/categories/{category}", actions.CategoriesIndex).Methods("GET")
	
	router.HandleFunc("/orders", actions.OrdersCreate).Methods("POST")
	router.HandleFunc("/orders/{order}/item", actions.OrdersUpdate).Methods("PUT")

	mongodb := dao.ShopDAO{Server: "mongodb://mongo:mongo@db:27017", Database: "shop"}
	err := mongodb.Connect()
	if err != nil {
		fmt.Println("Error connecting to mongodb:", err)
	}
	//r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	//r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	//r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	//r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	//r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}