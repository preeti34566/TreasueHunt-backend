package main

import (
	// "fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const port = "8080"

var atlasConnectionUri string

type application struct {
	Domain string
}

func main() {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	atlasConnectionUri = os.Getenv("MONGO_DB")

	var app application

	app.Domain = "treasureh0nt.onrender.com"

	log.Print("Starting application....")

	// starts a web server

	// serverErr := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())
	serverErr := http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), app.routes())

	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
