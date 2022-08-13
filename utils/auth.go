package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jadwalin/constants"
	"jadwalin/models"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{Error: "User unauthorized"})
	}
}
