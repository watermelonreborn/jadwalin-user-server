package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"jadwalin/constants"
	"jadwalin/models"
	"jadwalin/services"
)

func GetCode(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.JSON(http.StatusUnauthorized, models.Response{
			Error: "Unauthorized: Please provide a valid Bearer token in Authorization header",
		})
		return
	}

	rdb := services.RedisClient

	uuid := c.GetString(constants.UserIDKey)
	code := "example"

	// TODO: Check of code already exists in redis

	res := rdb.Set(context.Background(), code, uuid, constants.CodeExpirationDuration)
	if res.Err() != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Internal server error: Failed to generate new code",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: code})
}

func PostCode(c *gin.Context) {
	rdb := services.RedisClient

	var input models.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	res := rdb.Get(context.Background(), input.Code)
	if res.Err() != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Error: "Internal server error: Code has expired or is invalid",
		})
		return
	}

	// TODO: Register user to server
	fmt.Println(input)

	c.JSON(http.StatusOK, models.Response{Data: res.Val()})
}
