package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/errs"
	"gorm.io/datatypes"
)

type _completedHandler struct{}

// completedResponse a struct for storing the quiz's information and the responseAnswer struct
type completedResponse struct {
	Completions []*responseAnswer `json:"answers,omitempty"`
	Quiz        *models.Quiz      `json:"quiz"`
}

// responseAnswer stores a question, it answers possibilities, the correct answer and the answer by the user
type responseAnswer struct {
	Question   string         `json:"question"`
	Answers    datatypes.JSON `json:"answers"`
	Answer     string         `json:"answer"`
	UserAnswer string         `json:"user_answer"`
}

var CompletedHandler _completedHandler

// PaginateAll paginates the submitted quizzes
func (ch _completedHandler) PaginateAll(c *fiber.Ctx) error {
	rawPage := c.Query("page")
	pageNum, _ := strconv.Atoi(rawPage)

	if pageNum == 0 {
		pageNum++
	}

	authUser := GetAuthUser(c)

	completedQuizzes := repository.UserAnswer.Paginate(
		authUser.Id,
		5,
		pageNum,
		uint(0), // this means that no quiz filter applied, all returned
	)

	var completedRes []*completedResponse

	for _, cq := range completedQuizzes {
		completedRes = append(completedRes, ch.createCompletedResponse(cq))
	}

	return c.JSON(completedRes)
}

// FindOne only returns the specific submitted quiz by the "quizId" route param
func (ch _completedHandler) FindOne(c *fiber.Ctx) error {
	quizId, _ := strconv.Atoi(c.Params("quizId"))

	completedQuiz := repository.UserAnswer.Paginate(
		GetAuthUser(c).Id,
		1,
		1,
		uint(quizId), // this means that
	)

	if len(completedQuiz) != 1 {
		return errs.NotFound(c, errors.New("this quiz completion could not be found"))
	}

	return c.JSON(ch.createCompletedResponse(completedQuiz[0]))
}

// createCompletedResponse converts the model data to a proper response
func (ch _completedHandler) createCompletedResponse(u *models.UserAnswer) *completedResponse {
	// looping through the quizzes
	var completedRes completedResponse
	completedRes.Completions = ch.createResponseAnswers(u)

	u.HostedQuiz.Quiz.Questions = nil
	u.HostedQuiz.Quiz.User.Email = "<hidden>" // for safety
	completedRes.Quiz = u.HostedQuiz.Quiz
	return &completedRes
}

// createResponseAnswers converts the model data into an easier response for the client side
func (_completedHandler) createResponseAnswers(u *models.UserAnswer) []*responseAnswer {
	var resAnswers []*responseAnswer

	for _, answer := range u.Answers {
		for _, question := range u.HostedQuiz.Quiz.Questions {
			if question.Id == answer.QuestionId {
				resAnswers = append(resAnswers, &responseAnswer{
					Question:   question.Question,
					Answers:    question.Answers,
					Answer:     question.Answer,
					UserAnswer: answer.Answer,
				})
				break
			}
		}
	}

	return resAnswers
}
