package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/helpers"
	"github.com/goquiz/api/http/middleware"
	"github.com/goquiz/api/routes"
	"log"
)

func main() {
	app := fiber.New()
	helpers.Env.Load()
	database.Connect()
	app.Use(middleware.NewCors())

	routes.ApiV1.Add(app)

	app.Use(routes.NewNotFoundPage)

	log.Fatal(app.Listen(helpers.Env.ServerPort))
}
