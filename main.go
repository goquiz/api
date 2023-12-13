package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/helpers"
	"github.com/goquiz/api/routes"
	"log"
)

func main() {
	app := fiber.New()
	helpers.Env.Load()
	database.Connect()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Authorization,Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,X-Frontend-Client",
		AllowOrigins:     helpers.Env.Cors.AllowOrigins,
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.Api.Add(app)

	app.Use(routes.NewNotFoundPage)

	log.Fatal(app.Listen(helpers.Env.ServerPort))
}
