package main

import (
	"TreasureHunt-backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		ServerStatus string `json:"serverStatus"`
		DbStatus     bool   `json:"dbStatus"`
		Message      string `json:"message"`
		Version      string `json:"version"`
	}{
		ServerStatus: "active",
		Message:      "Treasure Hunt API",
		Version:      "1.1.0",
	}

	db, err := GetDb()

	dbErr := db.Client().Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(dbErr)
	} else {
		fmt.Println("Database available")
		payload.DbStatus = true
	}

	out, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) addUserDetail(w http.ResponseWriter, r *http.Request) {

	var user models.Users

	_ = json.NewDecoder(r.Body).Decode(&user)

	db, err := GetDb()

	if err != nil {
		log.Fatal(err)
	}

	collection := db.Collection("userDetails")

	// insert user data

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(result.InsertedID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func (app *application) getUserByUserId(w http.ResponseWriter, r *http.Request) {

	// userId := mux.Vars(r)["userId"]

	userId := chi.URLParam(r, "userId")

	db, err := GetDb()

	if err != nil {
		log.Fatal(err)
	}

	collection := db.Collection("userDetails")

	// find user detail with userId
	var user models.Users
	err = collection.FindOne(context.TODO(), bson.M{"userId": userId}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
