package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"

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

	scheduler := gocron.NewScheduler(time.UTC)
	// scheduler.Every(1).Day().At("01:00;13:00").Do(func() {
	scheduler.Every(1).Day().At("02:35;14:35").Do(func() {
		log.Println("[INFO] Scheduler started")
		services.SendDailyNotification()
		log.Println("[INFO] Scheduler finished")
	})

	scheduler.StartAsync()

	if err := s.ListenAndServe(); err != nil {
		panic("[ERROR] Failed to listen and serve")
	}
}
