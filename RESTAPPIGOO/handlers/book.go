package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/database"
	"github.com/vargasarielhernan/go-ejemplos.git/RESTAPPIGOO/models"
	"go.mongodb.org/mongo-driver/bson"
)

type bookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	ISBN      string `json:"isbn" bson:"isbn"`
	LibraryID string `json:"libraryId" bson:"libraryId"`
}

func CreateBook(c *fiber.Ctx) error {
	createData := new(bookDTO)
	if err := c.BodyParser(createData); err != nil {
		return err
	}
	//get collection
	coll := database.GetCollection("libraries")
	// get filter
	filter := bson.M{"id": createData.LibraryID}
	nBookData := models.Book{
		Title:  createData.Title,
		Author: createData.Author,
		ISBN:   createData.ISBN,
	}
	updatepayload := bson.M{"$push": bson.M{"books": nBookData}}
	//updateLibrary
	_, err := coll.UpdateOne(context.TODO(), filter, updatepayload)
	if err != nil {
		return err
	}

	return c.SendString("book created successfully")
}
