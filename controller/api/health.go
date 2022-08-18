package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"jadwalin/models"
	"jadwalin/services"
)

func HealthCheck(c *gin.Context) {

	log.Println("[INFO] Health checking...")

	// MongoDB check
	mongoStatus, mongoMessage := services.MongoHealthCheck()
	log.Println(mongoStatus, mongoMessage)

	// TODO: Check redis
	redisStatus, RedisMessage := services.RedisHealthCheck()
	log.Println(redisStatus, RedisMessage)

	data := models.Health{MongoDB: mongoMessage, Redis: RedisMessage}
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: data})
}
