package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"jadwalin/constants"
)

func AuthMiddleware(c *gin.Context) {

	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	c.Set(constants.IsAuthenticatedKey, true)

	// TODO: verify token

	// TODO: replace "user_id" with actual user id
	// c.Set(constants.UserIDKey, token.UID)
	c.Set(constants.UserIDKey, "user_id")
	c.Next()
}
