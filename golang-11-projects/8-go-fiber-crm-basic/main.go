package main

import (
	"fmt"

	"github.com/Mdromi/Crash-Course/golang-11-projects/8-go-fiber-crm-basic/database"
	"github.com/Mdromi/Crash-Course/golang-11-projects/8-go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead/", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBcon, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to the database")
	database.DBcon.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	defer func() {
		if db, err := database.DBcon.DB(); err == nil {
			db.Close()
		}
	}()
	app.Listen(":3000")
}
