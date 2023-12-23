package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/authorized"
	"github.com/goquiz/api/http/errs"
	"strconv"
)

type _quizHandler struct{}

var QuizHandler _quizHandler

func (_quizHandler) All(c *fiber.Ctx) error {
	quizzes := repository.Quiz.AllForUser(authorized.Authorized.User.Id)
	return c.JSON(fiber.Map{
		"quizzes": quizzes,
	})
}

func (q _quizHandler) WithQuestions(c *fiber.Ctx) error {
	quiz, err := q.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	return c.JSON(quiz)
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

func (_quizHandler) GetQuiz(c *fiber.Ctx) (*models.Quiz, error) {
	idInt, _ := strconv.Atoi(c.Params("id"))
	id := uint(idInt)
	quiz := repository.Quiz.WithQuestions(id, authorized.Authorized.User.Id)
	if quiz.Id == 0 {
		return nil, errors.New("this quiz could not be found or does not belongs to you")
	}
	return quiz, nil
}
