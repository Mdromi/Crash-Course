package app

import (
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/database"
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/handlers"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()

	mg := database.GetMongoInstance()
	employeeHandler := handlers.NewEmployeeHandler(mg)

	app.Get("/employee", employeeHandler.GetEmployees)
	app.Post("/employee", employeeHandler.CreateEmployee)
	app.Put("/employee/:id", employeeHandler.UpdateEmployee)
	app.Delete("/employee/:id", employeeHandler.DeleteEmployee)

	return app
}
