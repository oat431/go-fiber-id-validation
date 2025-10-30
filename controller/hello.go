package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oat431/go-fiber-id-validation/service"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.JSON(service.HelloWorldFunc())
}
