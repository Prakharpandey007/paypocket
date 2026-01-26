package app

import (
	"github.com/Prakharpandey007/paypocket/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.Container) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("/signup", h.UserHandler.Signup)
	users.Post("/signup-with-user", h.UserHandler.SignupReturnUser)
	users.Post("/login", h.UserHandler.Login)
	users.Get("/listuser", h.UserHandler.ListUsers)
}
