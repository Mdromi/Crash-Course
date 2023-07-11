package main

import (
	"log"

	"github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/app"
	_ "github.com/gofiber/fiber/v2"
)

func main() {
	app := app.NewApp()
	log.Fatal(app.Listen(":3000"))
}
