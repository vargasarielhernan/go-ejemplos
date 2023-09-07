package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/database"
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
