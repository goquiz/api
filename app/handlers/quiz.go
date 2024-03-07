package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/errs"
)

type _quizHandler struct{}

var QuizHandler _quizHandler

// All returns all the quizzes for the authenticated user
func (_quizHandler) All(c *fiber.Ctx) error {
	quizzes := repository.Quiz.AllForUser(GetAuthUser(c).Id)
	return c.JSON(fiber.Map{
		"quizzes": quizzes,
	})
}

// WithQuestions returns a quiz with all the questions for the authenticated user
func (q _quizHandler) WithQuestions(c *fiber.Ctx) error {
	quiz, err := q.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	return c.JSON(quiz)
}

// Create creates a quiz for the authenticated user
func (_quizHandler) Create(c *fiber.Ctx) error {
	authUserId := GetAuthUser(c).Id
	quizCount := repository.Quiz.CountForUser(authUserId)

	if quizCount >= 10 {
		return errs.BadRequest(c, errors.New("cannot create more than 10 quiz"))
	}

	quizRequest := requests.QuizValidation
	quiz := models.Quiz{
		Name:   quizRequest.Name,
		UserId: authUserId,
	}
	database.Database.Model(&models.Quiz{}).Create(&quiz)
	return c.JSON(fiber.Map{
		"message": "Your quiz has been created",
		"quiz_id": quiz.Id,
	})
}

// Update updates a quiz for the authenticated user
func (q _quizHandler) Update(c *fiber.Ctx) error {
	quiz, err := q.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	quizRequest := requests.QuizValidation
	quiz.Name = quizRequest.Name

	database.Database.Save(&quiz)
	return c.JSON(fiber.Map{
		"message": "Successfully modified",
	})
}

// Destroy destroys a quiz for the authenticated user
func (q _quizHandler) Destroy(c *fiber.Ctx) error {
	quiz, err := q.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	database.Database.Delete(&quiz)
	return c.JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}

// GetQuiz returns a quiz by the route param "id", globally
func (_quizHandler) GetQuiz(c *fiber.Ctx) (*models.Quiz, error) {
	idInt, _ := strconv.Atoi(c.Params("id"))
	id := uint(idInt)
	quiz := repository.Quiz.WithQuestions(id, GetAuthUser(c).Id)
	if quiz.Id == 0 {
		return nil, errors.New("this quiz could not be found or does not belongs to you")
	}
	return quiz, nil
}
