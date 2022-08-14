package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"jadwalin/config"
	"jadwalin/router"
	"jadwalin/services"
)

func main() {
	config.InitializeConfig()

	if config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	services.ConnectDB()
	services.InitializeRedis()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic("[ERROR] Failed to listen and serve")
	}
}
