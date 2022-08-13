package services

import (
	"github.com/go-redis/redis/v9"

	"jadwalin/config"
)

var RedisClient *redis.Client

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})
}
