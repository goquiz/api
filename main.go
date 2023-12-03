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
	database.Connect()
	helpers.Env.Load()

	routes.Api.Add(app)

	log.Fatal(app.Listen(helpers.Env.ServerPort))
}
