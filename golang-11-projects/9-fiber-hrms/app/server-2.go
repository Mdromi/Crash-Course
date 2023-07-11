package app

import (
	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/handlers"
	"github.com/gofiber/fiber/v2"
)

func newapp() *fiber.App {
	app := fiber.New()

	app.Get("/employee", handlers.GetEmployees)
	app.Post("/employee", handlers.CreateEmployee)
	app.Put("/employee/:id", handlers.UpdateEmployee)
	app.Delete("/employee/:id", handlers.DeleteEmployee)

	return app
}
