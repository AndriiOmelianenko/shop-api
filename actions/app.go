package actions

import (
	"log"
	"net/http"

	"github.com/AndriiOmelianenko/shop-api/dao"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

func Serve(c *cli.Context) {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")

	router.HandleFunc("/items", ItemsList).Methods("GET")
	router.HandleFunc("/items/{item}", ItemsIndex).Methods("GET")

	router.HandleFunc("/categories", CategoriesList).Methods("GET")
	router.HandleFunc("/categories/{category}", CategoriesIndex).Methods("GET")

	router.HandleFunc("/orders", OrdersCreate).Methods("POST")
	router.HandleFunc("/orders/{order}/item", OrdersUpdate).Methods("PUT")
	if dao.DB == nil {
		mongodb := dao.ShopDAO{Server: c.GlobalString("mongo"), Database: c.GlobalString("dbname")}
		mongodb.Connect()
	}
	//mongodb := dao.ShopDAO{Server: c.GlobalString("mongo"), Database: c.GlobalString("dbname")}
	//mongodb.Connect()

	//r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	//r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	//r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	//r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	//r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
