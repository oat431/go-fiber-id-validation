package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oat431/go-fiber-id-validation/config"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "hello world"})
	})

	port := config.GetEnv("PORT")

	app.Listen(":" + port)
}
