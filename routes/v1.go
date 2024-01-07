package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/handlers"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/helpers"
	"github.com/goquiz/api/http/middleware"
	"github.com/goquiz/api/http/sessions"
)

type _apiV1 struct{}

var ApiV1 _apiV1

func (_apiV1) Add(app *fiber.App) {
	// api group
	api := app.Group("/v1", sessions.NewGlobalSessionHandler)

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"hcaptcha_key": helpers.Env.HCaptcha.SiteKey,
		})
	})

	api.Post("/login", middleware.HCaptcha.New, requests.LoginRequest, handlers.AuthHandler.Login)
	api.Post("/register", middleware.HCaptcha.New, requests.RegisterRequest, handlers.AuthHandler.Register)
	api.Get("/logout", handlers.AuthHandler.Logout)
	api.Get("/email-verification/:token", handlers.AuthHandler.VerifyEmailAddress)
	api.Post("/reset-password/request", requests.RequestNewPasswordRequest, handlers.AuthHandler.RequestNewPassword)
	api.Post("/reset-password/change/:token", requests.ResetPasswordRequest, handlers.AuthHandler.ChangePassword)

	// authenticated group
	auth := api.Group("/", middleware.AuthMiddleware.Auth)
	auth.Get("/me", handlers.MeHandler.Hello)
	// quizzes
	auth.Post("/quiz", requests.QuizRequest, handlers.QuizHandler.Create)
	auth.Get("/quiz", handlers.QuizHandler.All)
	auth.Get("/quiz/:id<int>", handlers.QuizHandler.WithQuestions)
	auth.Put("/quiz/:id<int>", requests.QuizRequest, handlers.QuizHandler.Update)
	auth.Delete("/quiz/:id<int>", handlers.QuizHandler.Destroy)
	// adding or changing questions
	auth.Post("/quiz/:id<int>/questions", requests.QuestionRequest, handlers.QuestionHandler.Create)
	auth.Put("/quiz/:id<int>/questions/:questionId<int>", requests.QuestionRequest, handlers.QuestionHandler.Update)
	auth.Delete("/quiz/:id<int>/questions/:questionId<int>", handlers.QuestionHandler.Destroy)
	//
	auth.Post("/quiz/:id<int>/hosts", requests.HostRequest, handlers.HostHandler.New)
	auth.Put("/quiz/:id<int>/hosts/:hostId<int>/activity", handlers.HostHandler.ChangeActive)
	auth.Delete("/quiz/:id<int>/hosts/:hostId<int>", handlers.HostHandler.Destroy)
	auth.Get("/quiz/:id<int>/hosts", handlers.HostHandler.All)
	// Play routes
	auth.Get("/play/:public_key<maxLen(8)>/info", handlers.PlayHandler.Info)
	auth.Get("/play/:public_key<maxLen(8)>", handlers.PlayHandler.Play)
	auth.Post("/play/:public_key<maxLen(8)>", middleware.HCaptcha.New, requests.PlayRequest, handlers.PlayHandler.Submit)
	// Completed quizzes
	auth.Get("/completed", handlers.CompletedHandler.PaginateAll)
	auth.Get("/completed/:quizId<int>", handlers.CompletedHandler.FindOne)
	// Completed for host
	auth.Get("/completed/host/:hostId<int>", handlers.HostCompletionsHandler.Paginate)
}
