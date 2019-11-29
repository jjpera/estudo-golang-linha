package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "./go"
)

var dao = sw.LinhaRepository{}
var config = sw.Config{}

func init() {
	config.Read()

	log.Printf("Conectando no banco %s, %s", config.Server, config.Database)

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8082", router))
}
