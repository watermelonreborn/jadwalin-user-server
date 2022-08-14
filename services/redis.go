package services

import (
	"fmt"

	"github.com/go-redis/redis/v9"

	"jadwalin/config"
)

var RedisClient *redis.Client

func InitializeRedis() {
	fmt.Println("[INFO] Connecting to RedisDB")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	fmt.Println("[INFO] Connected to RedisDB")
}
