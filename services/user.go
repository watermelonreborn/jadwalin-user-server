package services

import (
	"context"
	"fmt"
	"jadwalin/constants"
	"jadwalin/converter"
	"jadwalin/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(userRegister models.UserRegister, authId string) (string, models.User) {
	log.Printf("%s CreateUser: %s, %s", constants.LogInfo, userRegister.DiscordID, userRegister.ServerID)
	userDb := MongoClient.Database("jadwalin").Collection("users")

	user, _ := GetUserByDiscordIDAndServerID(userRegister.DiscordID, userRegister.ServerID)
	if user.DiscordID != "" {
		log.Printf("%s User already registered: %s, %s", constants.LogError, userRegister.DiscordID, userRegister.ServerID)
		return constants.AlreadyRegistered, user
	}

	newUser := converter.UserRegisterToUser(userRegister, authId)
	res, err := userDb.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, userRegister.DiscordID)
		return constants.Error, newUser
	}

	server := GetServer(userRegister.ServerID)
	if server.ID == "" {
		CreateServerWithUser(res.InsertedID, newUser)
	} else {
		server.Members[userRegister.DiscordID] = res.InsertedID
		UpdateServer(userRegister.ServerID, server.Members)
	}

	return constants.Registered, user
}

func GetUser(userId string) {
	// TODO: create logic to get user from database
	fmt.Println(userId)
}

func GetUserByDiscordIDAndServerID(discordID string, serverID string) (models.User, error) {
	log.Printf("%s GetUserByDiscordIDAndServerID: %s, %s", constants.LogInfo, discordID, serverID)
	db := MongoClient.Database("jadwalin").Collection("users")

	var result models.User
	filter := bson.D{primitive.E{Key: "discord_id", Value: discordID}, primitive.E{Key: "server_id", Value: serverID}}
	err := db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("%s %s: %s, %s", constants.LogError, err, discordID, serverID)
		return models.User{}, err
	}
	return result, nil
}

func GetUserByAuthID(authID string) (models.User, error) {
	log.Printf("%s GetUserByAuthID: %s", constants.LogInfo, authID)
	db := MongoClient.Database("jadwalin").Collection("users")

	var result models.User
	filter := bson.D{primitive.E{Key: "auth_id", Value: authID}}
	err := db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, authID)
	}

	return result, err
}

func DeleteUser(userId string) string {
	log.Printf("%s DeleteUser: %s", constants.LogInfo, userId)
	db := MongoClient.Database("jadwalin").Collection("users")

	filter := bson.D{primitive.E{Key: "_id", Value: userId}}
	_, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("%s %s: %s", constants.LogError, err, userId)
	}

	return constants.Success
}
