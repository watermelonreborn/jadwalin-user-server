package services

import (
	"context"
	"fmt"
	"jadwalin/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(userRequest models.UserRequest) {
	// TODO: create logic to write user to database and save user's server
	MongoClient.StartSession()
	db := MongoClient.Database("jadwalin").Collection("users")
	doc := bson.D{{"user_id", "abc-321"}, {"server_id", "abc-123"}, {"token", "3131kshfakhfsa"}}
	db.InsertOne(context.TODO(), doc)
	fmt.Println(userRequest)
}

func GetUser(userId string) {
	// TODO: create logic to get user from database
	fmt.Println(userId)
}

func DeleteUser(userId string) {
	// TODO: create logic to delete user from database
	fmt.Println(userId)
}
