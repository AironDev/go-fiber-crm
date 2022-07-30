package main

import (
	"github.com/airondev/go-fiber-crm/pkg/interface/http"
	"github.com/airondev/go-fiber-crm/pkg/persistence/sqlite"
	"github.com/gofiber/fiber"
)

type Config struct {
}

func main() {
	app := fiber.New()
	http.SetupRoutes(app)
	err := app.Listen("3000")
	if err != nil {
		return
	}

	sqlite.InitDatabase()
}
