package main

import (
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/routes"
	"github.com/bndrmrtn/goquiz_api/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	database.Connect()
	utils.Env.Load()

	routes.Api.Add(app)

	log.Fatal(app.Listen(utils.Env.ServerPort))
}
