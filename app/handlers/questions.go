package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/authorized"
	"github.com/goquiz/api/http/errs"
	"slices"
	"strconv"
)

type _questionHandler struct{}

var QuestionHandler _questionHandler

// Create a new Question
func (q _questionHandler) Create(c *fiber.Ctx) error {
	quiz, err := QuizHandler.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}

	if repository.Quiz.QuestionsCount(quiz.Id) >= 15 {
		return errs.BadRequest(c, errors.New("you already added all the allowed questions"))
	}

	answers, err := q.checkRequestAnswers(c)
	if err != nil {
		return errs.BadRequest(c, err)
	}

	question := models.Question{
		Question: requests.QuestionValidation.Question,
		Answer:   requests.QuestionValidation.Answer,
		Answers:  answers,
		QuizId:   quiz.Id,
	}
	database.Database.Create(&question)
	return c.JSON(fiber.Map{
		"message": "Question successfully added",
		// question is returned cuz we don't know the id of the question
		// otherwise we have to pull all the questions from the api
		"question": question,
	})
}

// Update a specific question
func (q _questionHandler) Update(c *fiber.Ctx) error {
	question, err := q.getQuestion(c)
	if err != nil {
		return errs.NotFound(c, err)
	}

	answers, err := q.checkRequestAnswers(c)
	if err != nil {
		return errs.BadRequest(c, err)
	}

	question.Question = requests.QuestionValidation.Question
	question.Answers = answers
	question.Answer = requests.QuestionValidation.Answer

	database.Database.Save(&question)

	return c.JSON(fiber.Map{
		"message": "Successfully modified",
		// question is not returned cuz we know the id in frontend
	})
}

// Destroy a specific question
func (q _questionHandler) Destroy(c *fiber.Ctx) error {
	question, err := q.getQuestion(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	database.Database.Delete(&question)
	return c.JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}

// getQuestion (get a specific question from url params)
func (_questionHandler) getQuestion(c *fiber.Ctx) (*models.Question, error) {
	idInt, _ := strconv.Atoi(c.Params("id"))
	quizId := uint(idInt)
	if !repository.Quiz.IsBelongsToUser(quizId, authorized.Authorized.User.Id) {
		return nil, errors.New("this quiz does not belong to you")
	}

	idInt, _ = strconv.Atoi(c.Params("questionId"))
	id := uint(idInt)
	question := repository.Question.ForQuiz(id, quizId)
	if question.Id == 0 {
		return nil, errors.New("this question could not be found or does not belongs to you")
	}
	return question, nil
}

// checkRequestAnswers (checks if the request answers are valid or not)
func (_questionHandler) checkRequestAnswers(c *fiber.Ctx) ([]byte, error) {
	if !slices.Contains(requests.QuestionValidation.Answers, requests.QuestionValidation.Answer) {
		return nil, errs.BadRequest(c, errors.New("the answer must be one of the answers"))
	}

	return json.Marshal(requests.QuestionValidation.Answers)
}
