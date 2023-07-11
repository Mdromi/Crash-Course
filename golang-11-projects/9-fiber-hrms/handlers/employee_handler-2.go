package handlers

import (
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/database"
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEmployees(c *fiber.Ctx) error {
	mg := database.GetMongoInstance()
	query := bson.D{}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []models.Employee
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
}

func CreateEmployee(c *fiber.Ctx) error {
	mg := database.GetMongoInstance()
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result, err := mg.Db.Collection("employees").InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result.InsertedID)
}

func UpdateEmployee(c *fiber.Ctx) error {
	mg := database.GetMongoInstance()
	id := c.Params("id")
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	filter := bson.M{"_id": docID}
	update := bson.M{"$set": employee}
	_, err = mg.Db.Collection("employees").UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(200)
}

func DeleteEmployee(c *fiber.Ctx) error {
	mg := database.GetMongoInstance()
	id := c.Params("id")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	filter := bson.M{"_id": docID}
	_, err = mg.Db.Collection("employees").DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(200)
}
