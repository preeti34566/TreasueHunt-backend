package models

type Users struct {
	UserId       string `json:"userId" bson:"userId"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"pswrdHash" bson:"pswrdHash"`
	ImageURL     string `json:"imageUrl" bson:"imageUrl"`
}
