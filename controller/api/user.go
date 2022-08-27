package api

import (
	"encoding/json"
	"fmt"
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
			Code:  http.StatusUnauthorized,
			Error: "Unauthorized: Please provide a valid Bearer token in Authorization header",
		})
		return
	}

	uuid := c.GetString(constants.UserIDKey)
	code, err := services.CreateCode(uuid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:  http.StatusInternalServerError,
			Error: "Internal server error: Failed to generate new code",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Data: code})
}

func PostCode(c *gin.Context) {
	var input models.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	userId, err := services.UseCode(input.Code)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:  http.StatusNotFound,
			Error: "Not found: Code has expired or is invalid",
		})
		return
	}

	res, user := services.CreateUser(input, userId)

	if res == constants.AlreadyRegistered {
		log.Printf("[ERROR] user already registered: %s", input)
		c.JSON(http.StatusBadRequest, models.Response{Error: constants.AlreadyRegistered + ": user already registered", Data: user})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Data: userId})
}

func GetUser(c *gin.Context) {
	// TODO: Get user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func DeleteUser(c *gin.Context) {
	// TODO: Delete user from database
	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func SyncCalendar(c *gin.Context) {
	var input models.UserSearch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	user, err := services.GetUserByDiscordIDAndServerID(input.DiscordID, input.ServerID)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:  http.StatusNotFound,
			Error: "Not found: User is not registered",
		})
		return
	}

	res, err := services.SyncCalendar(user.AuthID)

	if res.StatusCode != 200 || err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:  http.StatusInternalServerError,
			Error: "Internal server error: An unexpected error occured",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 200, Data: "OK"})
}

func GetEvents(c *gin.Context) {
	var input models.UserSearch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	user, err := services.GetUserByDiscordIDAndServerID(input.DiscordID, input.ServerID)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:  http.StatusNotFound,
			Error: "Not found: User is not registered",
		})
		return
	}

	res, err := services.GetEvents(user.AuthID)

	if res.StatusCode != 200 || err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:  http.StatusInternalServerError,
			Error: "Internal server error: An unexpected error occured",
		})
		return
	}

	defer res.Body.Close()
	result := models.Response{}
	json.NewDecoder(res.Body).Decode(&result)
	result.Code = http.StatusOK
	c.JSON(http.StatusOK, result)
}

func GetSummary(c *gin.Context) {
	var input models.UserSummary
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	user, err := services.GetUserByDiscordIDAndServerID(input.DiscordID, input.ServerID)

	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Code:  http.StatusNotFound,
			Error: "Not found: User is not registered",
		})
		return
	}

	res, err := services.GetSummary(user.AuthID, input.Days, input.StartHour, input.EndHour)

	if res.StatusCode != 200 || err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:  http.StatusInternalServerError,
			Error: "Internal server error: An unexpected error occured",
		})
		return
	}

	defer res.Body.Close()
	result := models.Response{}
	json.NewDecoder(res.Body).Decode(&result)
	result.Code = http.StatusOK
	c.JSON(http.StatusOK, result)
}
