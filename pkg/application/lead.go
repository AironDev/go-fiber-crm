package application

import (
	"github.com/airondev/go-fiber-crm/pkg/persistence/sqlite"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int64  `json:"phone"`
}

var (
	database = *sqlite.DbConn
)

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	database.Find(&leads)
	err := c.JSON(leads)
	return err
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.Find(&lead, id)
	err := c.JSON(lead)
	return err
}

func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)
	err := c.BodyParser(lead)
	if err != nil {
		c.Status(503).Send([]byte("Error occurred"))
	}
	database.Create(&lead)
	c.JSON(lead)
	return err
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	if database.Find(&lead, id); lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with Id"))
		return nil
	}
	database.Delete(&lead)
	c.Status(500).Send([]byte("Lead deleted"))
	return nil
}
