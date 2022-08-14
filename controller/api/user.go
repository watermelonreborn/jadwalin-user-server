package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jadwalin/models"
	"jadwalin/services"
)

func CreateUser(c *gin.Context) {
	// TODO: Write user to database

	var request models.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	services.CreateUser(request)

	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func GetUser(c *gin.Context) {
	// TODO: Get user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func DeleteUser(c *gin.Context) {
	// TODO: Delete user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}
