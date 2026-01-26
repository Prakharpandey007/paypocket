package app

import (
	"log"

	"github.com/Prakharpandey007/paypocket/config"
	"github.com/Prakharpandey007/paypocket/internal/db"
	"github.com/Prakharpandey007/paypocket/internal/handler"
	"github.com/Prakharpandey007/paypocket/internal/model"
	"github.com/Prakharpandey007/paypocket/internal/repository"
	"github.com/Prakharpandey007/paypocket/internal/service"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	cfg := config.Load()

	db.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	app := fiber.New()

	// 🔹 Repository
	userRepo := repository.NewRepository[model.User](db.DB)

	// 🔹 Service
	userService := service.NewUserService(userRepo)

	// 🔹 Handlers (DI happens here)
	handlers := handler.NewContainer(userService)

	// 🔹 Routes
	SetupRoutes(app, handlers)

	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
