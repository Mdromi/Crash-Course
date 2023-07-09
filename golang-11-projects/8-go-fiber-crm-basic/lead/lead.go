package lead

import (
	"github.com/Mdromi/Crash-Course/golang-11-projects/8-go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBcon
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBcon
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBcon
	newLead := new(Lead)
	if err := c.BodyParser(newLead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&newLead)
	return c.JSON(newLead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBcon

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
