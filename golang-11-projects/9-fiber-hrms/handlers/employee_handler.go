package handlers

import (
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/database"
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeHandler struct {
	MongoInstance *database.MongoInstance
}

func NewEmployeeHandler(mg *database.MongoInstance) *EmployeeHandler {
	return &EmployeeHandler{
		MongoInstance: mg,
	}
}

func (h *EmployeeHandler) GetEmployees(c *fiber.Ctx) error {
	query := bson.D{}
	cursor, err := h.MongoInstance.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []models.Employee
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
}

func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	result, err := h.MongoInstance.Db.Collection("employees").InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result.InsertedID)
}

func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
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
	_, err = h.MongoInstance.Db.Collection("employees").UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(200)
}

func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	filter := bson.M{"_id": docID}
	_, err = h.MongoInstance.Db.Collection("employees").DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(200)
}
