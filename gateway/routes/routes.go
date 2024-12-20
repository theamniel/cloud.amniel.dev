package routes

import (
	"cloud.amniel.dev/gateway/controllers"
	"github.com/gofiber/fiber/v2"
)

func Configure(app *fiber.App) {
	file := app.Group("/f")
	file.Get("/", controllers.GetAllFile())
	file.Get("/:id", controllers.GetFile())
	file.Post("/upload", controllers.AddFile())
	file.Delete("/:id", controllers.DeleteFile())

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"message": "Endpoint is not found"})
	})
}
