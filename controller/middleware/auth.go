package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"jadwalin/config"
	"jadwalin/constants"
	"jadwalin/models"
)

func AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", config.AppConfig.AuthURL+"/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+idToken)
	res, err := client.Do(req)

	if res.StatusCode != 200 || err != nil {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	defer res.Body.Close()
	result := models.Response{}
	json.NewDecoder(res.Body).Decode(&result)
	data, _ := result.Data.(map[string]interface{})

	if data["id"] == nil {
		c.Set(constants.IsAuthenticatedKey, false)
		c.Next()
		return
	}

	c.Set(constants.IsAuthenticatedKey, true)
	c.Set(constants.UserIDKey, data["id"])
	c.Next()
}
