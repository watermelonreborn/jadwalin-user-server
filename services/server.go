package services

import (
	"context"
	"jadwalin/constants"
	"jadwalin/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateServer(uid interface{}, user models.User) (string, error) {
	log.Printf("%s CreateServer: %s, %s", constants.LogInfo, user.DiscordID, user.ServerID)
	db := MongoClient.Database("jadwalin").Collection("servers")

	members := make(map[string]interface{})
	members[user.DiscordID] = uid
	server := models.Server{
		ID:        user.ServerID,
		Members:   members,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := db.InsertOne(context.TODO(), server)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, user.ServerID)
		return constants.Error, err
	}

	return constants.Success, err
}

func GetServer(serverId string) models.Server {
	log.Printf("%s GetServer: %s", constants.LogInfo, serverId)
	db := MongoClient.Database("jadwalin").Collection("servers")

	var result models.Server
	filter := bson.D{primitive.E{Key: "_id", Value: serverId}}
	err := db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, serverId)
	}

	return result
}

func UpdateServer(serverId string, serverMembers map[string]interface{}) (string, error) {
	log.Printf("%s UpdateServer: %s", constants.LogInfo, serverId)
	db := MongoClient.Database("jadwalin").Collection("servers")

	filter := bson.D{primitive.E{Key: "_id", Value: serverId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "members", Value: serverMembers}}}}
	_, err := db.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, serverId)
		return constants.Error, err
	}

	return constants.Success, err
}
