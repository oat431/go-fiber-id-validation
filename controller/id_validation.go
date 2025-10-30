package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oat431/go-fiber-id-validation/service"
)

func ValidateNID(c *fiber.Ctx) error {
	return c.JSON(service.ValidateNID(string(c.Query("id"))))
}
