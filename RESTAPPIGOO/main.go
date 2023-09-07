package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/database"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/handlers"
)

func main() {
	//init app
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseMongoDB()

	app := GenerateApp()

	//listen port
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
func initApp() error {
	err := loadEnv()
	if err != nil {
		return err
	}
	err = database.StarMongoDB()
	if err != nil {
		return err
	}
	return nil
}
func GenerateApp() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})
	//create routes
	libGroup := app.Group("library")
	libGroup.Get("/", handlers.GETLibraries)
	libGroup.Post("/", handlers.CreateLibrary)
	return app
}

func loadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
