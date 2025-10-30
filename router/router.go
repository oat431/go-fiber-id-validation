package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/oat431/go-fiber-id-validation/config"
	"github.com/oat431/go-fiber-id-validation/controller"
)

func init() {
	log.Println("Creating router")
}

func StartServer() {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	hello := v1.Group("/hello")
	hello.Get("/", controller.HelloWorld)

	port := ":" + config.GetEnv("PORT")
	app.Listen(port)
}
