package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jadwalin/models"
	"jadwalin/services"
)

func GetServer(c *gin.Context) {
	id := c.Param("serverid")
	server := services.GetServer(id)

	if server.ID == "" {
		c.JSON(http.StatusNotFound, models.Response{
			Error: "Not found: Server doesn't exists.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: server})
}

func CreateServer(c *gin.Context) {
	var input models.ServerCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	server := services.GetServer(input.ServerID)

	if server.ID != "" {
		c.JSON(http.StatusConflict, models.Response{
			Error: "Conflict: Server already exists.",
		})
		return
	}

	_, err := services.CreateServer(input.ServerID, input.TextChannel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Internal server error: An unexpected error occured.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: "OK"})
}

func UpdateServer(c *gin.Context) {
	var input models.ServerCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	_, err := services.UpdateTextChannel(input.ServerID, input.TextChannel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Internal server error: An unexpected error occured.",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{Data: "OK"})
}
