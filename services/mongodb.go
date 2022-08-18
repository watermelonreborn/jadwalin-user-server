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

func ConnectDB() (*mongo.Client, error) {
	log.Println("[INFO] Connecting to MongoDB")

	credential := options.Credential{
		Username: config.AppConfig.DBUsername,
		Password: config.AppConfig.DBPassword,
	}

	dbUri := config.AppConfig.DBHost + ":" + config.AppConfig.DBPort + "/"
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUri)

	if credential.Username != "" && credential.Password != "" {
		clientOptions.SetAuth(credential)
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Printf("[ERROR] Error creating new mongo client: %s", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("[ERROR] Error connecting to mongo client: %s", err)
		return nil, err
	}

	MongoClient = client

	client.Ping(context.TODO(), nil)

	log.Println("[INFO] Connected to MongoDB")

	return client, nil
}

func MongoHealthCheck() (string, string) {
	log.Println("[INFO] MongoDB checking...")

	err := MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return "[ERROR]", err.Error()
	}

	return "[INFO]", "MongoDB is connected!"
}
