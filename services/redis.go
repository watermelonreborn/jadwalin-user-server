package services

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"

	"jadwalin/config"
)

var RedisClient *redis.Client

func InitializeRedis() {
	log.Println("[INFO] Connecting to RedisDB")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	log.Println("[INFO] Connected to RedisDB")
}

func RedisHealthCheck() (string, string) {
	log.Println("[INFO] Redis checking...")

	err := RedisClient.Ping(context.TODO())
	if err.Err() != nil {
		return "[ERROR]", err.Err().Error()
	}

	return "[INFO]", "Redis is connected!"
}
