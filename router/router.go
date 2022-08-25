package router

import (
	"github.com/gin-gonic/gin"

	"jadwalin/controller/api"
	"jadwalin/controller/middleware"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(
		middleware.AuthMiddleware,
		middleware.CorsMiddleware,
	)

	apiRoute := router.Group("/api")
	{
		user := apiRoute.Group("/user")
		{
			user.GET("/code", api.GetCode)
			user.POST("/code", api.PostCode)
			user.POST("/sync", api.SyncCalendar)
		}
		health := apiRoute.Group("/health")
		{
			health.GET("", api.HealthCheck)
		}
		server := apiRoute.Group("/server")
		{
			server.GET("/:serverid", api.GetServer)
			server.POST("/create", api.CreateServer)
			server.POST("/update", api.UpdateServer)
		}
	}

	return router
}
