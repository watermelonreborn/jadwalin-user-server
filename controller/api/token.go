package api

import (
	"jadwalin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteToken(c *gin.Context) {
	// TODO: Write token to redis
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}
