package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
)

type ServiceConfig struct {
	DBConnection    string `json:"dbconnection"`
	RestfulEndpoint string `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBConnectionDefault,
		RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
