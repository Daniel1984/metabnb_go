package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"metabnb/controllers"
	"metabnb/lib/configuration"
	"metabnb/lib/persistence/mongolayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := mongolayer.NewMongoDBLayer(config.DBConnection)

	router := httprouter.New()
	router.GET("/", controllers.GetListings)
	log.Fatal(http.ListenAndServe(":8081", router))
}
