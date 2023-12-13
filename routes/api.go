package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/handlers"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/http/middleware"
	"github.com/goquiz/api/http/sessions"
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
	auth.Get("/quiz/:id<int>", handlers.QuizHandler.WithQuestions)
	// adding or changing questions
	auth.Post("/quiz/:id<int>/questions", requests.QuestionRequest, handlers.QuestionHandler.AddQuestions)
	auth.Post("/quiz/:id<int>/host", requests.HostRequest, handlers.HostHandler.New)
	auth.Put("/quiz/host/:id<int>/active", handlers.HostHandler.ChangeActive)
}
