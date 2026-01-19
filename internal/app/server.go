package app

import (
	"log"

	"github.com/Prakharpandey007/paypocket/config"
	"github.com/Prakharpandey007/paypocket/internal/db"
	"github.com/Prakharpandey007/paypocket/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	cfg := config.Load()
	db.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	// migrations.Migrate()

	app := fiber.New()

	handlers := handler.NewContainer()
	SetupRoutes(app, handlers)
	// InitCronJobs()

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
