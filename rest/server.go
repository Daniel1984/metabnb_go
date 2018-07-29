package rest

import (
	"flag"
	"fmt"
	"metabnb/lib/configuration"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"metabnb/controllers"
	"metabnb/lib/persistence/mongolayer"
)

func Server() error {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := mongolayer.NewMongoDBLayer(config.DBConnection)
	router := httprouter.New()
	router.GET("/", controllers.GetListings)
	return http.ListenAndServe(config.RestfulEndpoint, router)
}
