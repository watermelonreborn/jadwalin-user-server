package services

import (
	"context"
	"jadwalin/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func CloseDB(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}

	}()
}

func ConnectDB() (*mongo.Client, context.Context, context.CancelFunc, error) {

	log.Println("[INFO] Connecting to MongoDB")
	log.Println(config.AppConfig.DBHost)
	log.Println(config.AppConfig.DBPort)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	dbUri := config.AppConfig.DBHost + ":" + config.AppConfig.DBPort + "/"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri), options.Client().SetAuth(options.Credential{
		Username: config.AppConfig.DBUsername,
		Password: config.AppConfig.DBPassword,
	}))

	MongoClient = client

	MongoClient.Ping(context.TODO(), nil)

	log.Println("[INFO] Connected to MongoDB")

	return client, ctx, cancel, err
}

func MongoHealthCheck() (string, string) {

	log.Println("[INFO] MongoDB checking...")
	err := MongoClient.Ping(context.TODO(), nil)

	if err != nil {
		return "[ERROR]", "MongoDB is not available!"
	}

	return "[INFO]", "MongoDB is connected!"
}
