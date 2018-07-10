package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"metabnb/controllers"
)

func main() {
	router := httprouter.New()
	router.GET("/", controllers.GetListings)
	log.Fatal(http.ListenAndServe(":8081", router))
}
