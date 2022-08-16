package services

import (
	"context"
	"fmt"
	"jadwalin/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateServer(uid interface{}, user models.User) (string, error) {

	db := MongoClient.Database("jadwalin").Collection("servers")

	var foundServer models.Server

	err := db.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: user.ServerID}}).Decode(&foundServer)

	if err != nil {
		fmt.Println(err)
	}

	members := foundServer.Members

	members[user.DiscordID] = uid
	server := models.Server{
		ID:        user.ServerID,
		Members:   members,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.UpdateOne(context.TODO(), server.ID, server)

	return "success", err
}

func GetServer(serverId string) models.Server {
	db := MongoClient.Database("jadwalin").Collection("servers")

	var result models.Server
	filter := bson.D{primitive.E{Key: "_id", Value: serverId}}
	err := db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	return result
}

func UpdateServer(serverId string, server models.Server) {
	// TODO: Update server
}
