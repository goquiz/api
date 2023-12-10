package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/authorized"
)

type _quizHandler struct{}

var QuizHandler _quizHandler

func (_quizHandler) All(*fiber.Ctx) error {
	repository.Quiz.All()
	return nil
}

func (_quizHandler) Create(c *fiber.Ctx) error {
	quizRequest := requests.QuizValidation
	quiz := models.Quiz{
		Name:   quizRequest.Name,
		UserId: authorized.Authorized.User.Id,
	}
	database.Database.Model(&models.Quiz{}).Create(&quiz)
	return c.JSON(fiber.Map{
		"message": "Your quiz has been created",
		"quiz_id": quiz.Id,
	})
}
