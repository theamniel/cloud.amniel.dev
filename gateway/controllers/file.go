package controllers

import "github.com/gofiber/fiber/v2"

// GET /f
func GetAllFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}

// GET /f/{id}
func GetFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}

// POST /f/upload
func AddFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}

// DELETE /f/{id}
func DeleteFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}
}
