package handlers

import (
	"github.com/bndrmrtn/goquiz_api/app/repository"
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/database/models"
	"github.com/bndrmrtn/goquiz_api/http/authorized"
	"github.com/gofiber/fiber/v2"
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
