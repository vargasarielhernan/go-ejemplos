package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/database"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/models"
	"go.mongodb.org/mongo-driver/bson"
)

type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exist" bson:"no_exist"`
}

// GET library
func GETLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")
	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	var libraries []models.Library
	if err := cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}
	return c.JSON(libraries)
}

// POST
func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)
	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}
	nLibrary.Empty = make([]string, 0)
	libraryCollection := database.GetCollection("libraries")
	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"id": nDoc.InsertedID})
}

// delete by id
func DeleteLibrary(c *fiber.Ctx) error {
	id := c.Params("id")
	libraryCollection := database.GetCollection("libraries")
	_, err := libraryCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return c.SendString("delete successfully")
}
