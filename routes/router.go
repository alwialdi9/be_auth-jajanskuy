package routes

import (
	"github.com/alwialdi9/be_auth-jajanskuy/handlers"
	"github.com/alwialdi9/be_auth-jajanskuy/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() *fiber.App {
	// Initialize the router
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     10 * 1024 * 1024, // 10 MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		// AllowHeaders:  "Origin,Content-Type,Accept,Authorization",
		ExposeHeaders: "Content-Length,Content-Type",
	}))

	api := app.Group("/api")
	api.Use(middlewares.PathLogger)
	// Define the API routes
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	api.Post("/signup", handlers.SignUp)

	return app
}
