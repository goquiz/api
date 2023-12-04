package routes

import (
	"github.com/bndrmrtn/goquiz_api/app/handlers"
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/bndrmrtn/goquiz_api/http/middleware"
	"github.com/bndrmrtn/goquiz_api/http/sessions"
	"github.com/gofiber/fiber/v2"
)

type api struct{}

var Api api

func (api) Add(app *fiber.App) {
	api := app.Group("/api", sessions.NewGlobalSessionHandler)

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Api is UP!")
	})

	api.Post("/login", requests.LoginRequest, handlers.Auth.Login)
	api.Post("/register", requests.RegisterRequest, handlers.Auth.Register)

	auth := api
	auth.Use(middleware.AuthMiddleware.Auth)
	auth.Get("/me", handlers.Me.Hello)
}
