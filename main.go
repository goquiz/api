package main

import (
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/helpers"
	"github.com/bndrmrtn/goquiz_api/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	helpers.Env.Load()
	database.Connect()

	routes.Api.Add(app)
	app.Use(routes.NotFoundPage.New)

	log.Fatal(app.Listen(helpers.Env.ServerPort))
}
