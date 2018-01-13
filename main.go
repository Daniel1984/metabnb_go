package main

import (
	"log"
	"net/http"

	handlers "./handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	log.Fatal(http.ListenAndServe(":8081", router))
}
