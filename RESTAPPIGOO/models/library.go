package models

type Library struct {
	Id      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Books   []Book `json:"books" bson:"books"`
}
