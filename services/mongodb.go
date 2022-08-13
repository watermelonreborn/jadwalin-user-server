package services

import (
	"context"
	"fmt"
	"jadwalin/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CloseDB(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {

		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}

	}()
}

func ConnectDB() (*mongo.Client, context.Context, context.CancelFunc, error) {

	fmt.Println("[INFO] Connecting to MongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig.DBURL))

	fmt.Println("[INFO] Connected to MongoDB")

	return client, ctx, cancel, err
}
