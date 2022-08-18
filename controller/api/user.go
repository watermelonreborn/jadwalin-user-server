package api

import (
	"log"
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

	uuid := c.GetString(constants.UserIDKey)
	code, err := services.CreateCode(uuid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Internal server error: Failed to generate new code",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: code})
}

func PostCode(c *gin.Context) {
	var input models.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	res, err := services.UseCode(input.Code)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Error: "Internal server error: Code has expired or is invalid",
		})
		return
	}

	res, user := services.CreateUser(input)

	if res == constants.AlreadyRegistered {
		log.Printf("[ERROR] user already registered: %s", input)
		c.JSON(http.StatusBadRequest, models.Response{Error: constants.AlreadyRegistered + ": user already registered", Data: user})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: res})
}

func GetUser(c *gin.Context) {
	// TODO: Get user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func DeleteUser(c *gin.Context) {
	// TODO: Delete user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}
