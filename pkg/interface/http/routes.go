package http

import (
	"github.com/airondev/go-fiber-crm/pkg/application"
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", application.GetLeads)
	app.Get("/api/v1/leads/:id", application.GetLead)
	app.Post("/api/v1/leads", application.NewLead)
	app.Delete("/api/v1/leads/:id", application.DeleteLead)
}
