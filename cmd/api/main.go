package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New()

	app.Get("*", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"app": "auth-svc"})
	})

	// Respond all Options as OK
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
