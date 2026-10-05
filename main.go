package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"zoie/config"
	"zoie/db"
	"zoie/handlers"
)

func main() {
	cfg := config.Load()
	conn := db.Connect(cfg.DatabaseURL)

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Static assets
	r.Static("/static", "./static")

	// Handlers
	homeHandler := handlers.NewHomeHandler(cfg)
	donateHandler := handlers.NewDonateHandler(conn, cfg)

	// Routes
	r.GET("/", homeHandler.Show)
	r.POST("/donate/initiate", donateHandler.Initiate)
	r.GET("/donate/callback", donateHandler.Callback)
	r.POST("/webhooks/paystack", donateHandler.Webhook)

	log.Printf("Zoie Foundation running on :%s (%s)", cfg.Port, cfg.AppEnv)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
