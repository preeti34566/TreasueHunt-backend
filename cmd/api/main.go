package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

const atlasConnectionUri = "mongodb+srv://anubhav11697:myMongo123@myfirstcluster.hfdwigv.mongodb.net/?retryWrites=true&w=majority"

type application struct {
	Domain string
}

func main() {

	var app application

	app.Domain = "example.com"

	log.Print("Starting application....")

	// starts a web server

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
