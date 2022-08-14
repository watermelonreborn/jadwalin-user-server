package router

import (
	"github.com/gin-gonic/gin"

	"jadwalin/controller/api"
	"jadwalin/controller/middleware"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	apiRoute := router.Group("/api")
	apiRoute.Use(
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	)
	{
		user := apiRoute.Group("/user")
		{
			// userRoute.POST("", api.CreateUser)
			// route to get user with user id string
			// userRoute.GET("/:userId", api.GetUser)
			user.GET("/code", api.GetCode)
			user.POST("/code", api.PostCode)
		}
		health := apiRoute.Group("/health")
		{
			health.GET("", api.HealthCheck)
		}
	}

	return router
}
