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
	auth.Delete("/quiz/:id<int>", handlers.QuizHandler.Destroy)
	// adding or changing questions
	auth.Post("/quiz/:id<int>/questions", requests.QuestionRequest, handlers.QuestionHandler.Create)
	auth.Put("/quiz/:id<int>/questions/:questionId<int>", requests.QuestionRequest, handlers.QuestionHandler.Update)
	auth.Delete("/quiz/:id<int>/questions/:questionId<int>", handlers.QuestionHandler.Destroy)
	//
	auth.Post("/quiz/:id<int>/hosts", requests.HostRequest, handlers.HostHandler.New)
	auth.Put("/quiz/:id<int>/hosts/hostId<int>/active", handlers.HostHandler.ChangeActive)
	auth.Get("/quiz/:id<int>/hosts", handlers.HostHandler.All)
}
