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
)

type _playHandler struct{}

var PlayHandler _playHandler

// Info returns the information about the quiz
func (p _playHandler) Info(c *fiber.Ctx) error {
	hosted, err := p.getHostedQuiz(c.Params("public_key"))

	if err != nil {
		return errs.NotFound(c, err)
	}

	if err := p.canAuthUserPlay(hosted.QuizId); err != nil {
		return errs.BadRequest(c, err)
	}

	return c.JSON(fiber.Map{
		"quiz_name": hosted.Quiz.Name,
		"username":  hosted.Quiz.User.Username,
	})
}

// Play returns the quiz data to play
func (p _playHandler) Play(c *fiber.Ctx) error {
	hosted := repository.HostedQuiz.FindByPublicKey(c.Params("public_key"))
	if hosted.Id == 0 {
		return errs.NotFound(c, errors.New("couldn't find this hosted quiz"))
	}

	if err := p.canAuthUserPlay(hosted.Id); err != nil {
		return errs.BadRequest(c, err)
	}

	questions := repository.Question.ForPlayers(hosted.QuizId)

	return c.JSON(questions)
}

// Submit validates and saves the answers
func (p _playHandler) Submit(c *fiber.Ctx) error {
	hosted := repository.HostedQuiz.FindByPublicKey(c.Params("public_key"))
	if hosted.Id == 0 {
		return errs.NotFound(c, errors.New("couldn't find this hosted quiz"))
	}

	if err := p.canAuthUserPlay(hosted.Id); err != nil {
		return errs.BadRequest(c, err)
	}

	var answers []*models.Answer

	questions := repository.Question.ForPlayers(hosted.QuizId)

	if len(requests.PlayValidation.Answers) != len(questions) {
		return errs.BadRequest(c, errors.New("invalid number of answers"))
	}

	userAnswer := models.UserAnswer{
		HostedQuizId: hosted.Id,
		UserId:       authorized.Authorized.User.Id,
	}

	err := database.Database.Create(&userAnswer).Error

	if err != nil {
		return errs.InternalServerError(c, errors.New("invalid number of answers"))
	}

	for i, q := range questions {
		answers = append(answers, &models.Answer{
			UserAnswerId: userAnswer.Id,
			QuestionId:   q.Id,
			Answer:       requests.PlayValidation.Answers[i],
		})
	}

	result := database.Database.Create(answers)

	if result.Error != nil {
		return errs.InternalServerError(c, result.Error)
	}

	return c.JSON(fiber.Map{
		"message": "Successfully submitted",
		"quiz_id": hosted.QuizId,
	})
}

// getHostedQuiz returns a models.HostedQuiz by a given public key (or public identifier: 514 158)
func (_playHandler) getHostedQuiz(publicKey string) (*models.HostedQuiz, error) {
	hosted := repository.HostedQuiz.FindByPublicKeyWithQuizUser(publicKey)
	if hosted.Id == 0 {
		return nil, errors.New("couldn't find this hosted quiz")
	}
	return hosted, nil
}

// canAuthUserPlay returns an error if the authenticated user already played the quiz before
func (_playHandler) canAuthUserPlay(hostedQuizId uint) error {
	canPlay := !repository.UserAnswer.IsUserAlreadyPlayed(hostedQuizId, authorized.Authorized.User.Id)
	if canPlay == true {
		return nil
	}
	return errors.New("you already completed this quiz")
}
