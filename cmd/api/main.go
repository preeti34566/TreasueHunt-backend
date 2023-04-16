package main

import (
	// "fmt"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetDb() (*mongo.Database, error) {
	opts := options.Client().ApplyURI(atlasConnectionUri)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	} else {

		// check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("DB Connction succeeded!")
		}
	}

	db := client.Database("TreasureHunt")

	return db, nil
}
