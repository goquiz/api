package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/errs"
	"slices"
)

type _questionHandler struct{}

var QuestionHandler _questionHandler

func (q _questionHandler) AddQuestions(c *fiber.Ctx) error {
	quiz, err := QuizHandler.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}

	if !slices.Contains(requests.QuestionValidation.Answers, requests.QuestionValidation.Answer) {
		return errs.BadRequest(c, errors.New("the answer must be one of the answers"))
	}

	answers, err := json.Marshal(requests.QuestionValidation.Answers)

	if err != nil {
		return errs.BadRequest(c, err)
	}

	question := models.Question{
		Question: requests.QuestionValidation.Question,
		Image:    requests.QuestionValidation.Image,
		Answer:   requests.QuestionValidation.Answer,
		Answers:  answers,
		QuizId:   quiz.Id,
	}
	database.Database.Create(&question)
	return c.JSON(fiber.Map{
		"message": "Question successfully added",
	})
}
