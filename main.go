package main

import (
	"log"
	"net/http"

	controllers "./controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", controllers.Listing)
	log.Fatal(http.ListenAndServe(":8081", router))
}
