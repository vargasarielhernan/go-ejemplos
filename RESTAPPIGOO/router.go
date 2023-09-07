package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/handlers"
)

func GenerateApp() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	//create routes
	libGroup := app.Group("library")
	libGroup.Get("/", handlers.TestHandler)
	return app
}
