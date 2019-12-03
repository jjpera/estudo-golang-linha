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
	rp "./go/repository"
	sv "./go/service"
)

var repository = rp.LinhaRepository{}
var config = sw.Config{}
var service = sv.HistoricoService{}

func init() {
	config.Read()

	log.Printf("Conectando no banco %s, %s", config.Server, config.Database)

	repository.Server = config.Server
	repository.Database = config.Database
	repository.Connect()

	service.SetHistoricourl(config.Historicourl)
}

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8082", router))
}
