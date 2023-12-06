package routes

import (
	"github.com/bndrmrtn/goquiz_api/app/handlers"
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/bndrmrtn/goquiz_api/http/middleware"
	"github.com/bndrmrtn/goquiz_api/http/sessions"
	"github.com/gofiber/fiber/v2"
)

type _api struct{}

var Api _api

func (_api) Add(app *fiber.App) {
	// api group
	api := app.Group("/api", sessions.NewGlobalSessionHandler)

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Api is UP!")
	})

	api.Post("/login", requests.LoginRequest, handlers.AuthHandler.Login)
	api.Post("/register", requests.RegisterRequest, handlers.AuthHandler.Register)
	api.Get("/logout", handlers.AuthHandler.Logout)

	// authenticated group
	auth := api.Group("/", middleware.AuthMiddleware.Auth)
	auth.Get("/me", handlers.MeHandler.Hello)
	auth.Post("/quiz", requests.QuizRequest, handlers.QuizHandler.Create)
	auth.Get("/quiz", handlers.QuizHandler.All)
}
