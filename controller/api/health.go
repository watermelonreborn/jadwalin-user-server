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
	status, message := services.MongoHealthCheck()
	log.Println(status, message)

	// TODO: Check redis

	data := models.Health{MongoDB: message, Redis: "Redis not presence!"}
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: data})
}
