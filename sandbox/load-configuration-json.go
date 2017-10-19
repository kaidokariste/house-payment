package main

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

// MyConfig holds the configuration values from config.json file
var MyConfig configuration

// Initialize MyConfig
func initConfig() {
	loadMyConfig()
}
// Reads config.json and decode into MyConfig
func loadMyConfig() {
	file, err := os.Open("common/config.example.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	// NewDecoder returns a new Decoder struct that reads from file.
	decoder := json.NewDecoder(file)
	MyConfig = configuration{}

	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by MyConfig
	err = decoder.Decode(&MyConfig)
	if err != nil {
		log.Fatalf("[loadMyConfig]: %s\n", err)
	}
}

func main() {
	// Initialise load
	initConfig()
	// Print out one value from example configuration file
	fmt.Println("Server in example configuration file", MyConfig.Server)

}