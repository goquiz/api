package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/http/errs"
)

type _playHandler struct{}

var PlayHandler _playHandler

func (p _playHandler) Info(c *fiber.Ctx) error {
	hosted, err := p.getHostedQuiz(c.Params("public_key"))
	if err != nil {
		return errs.NotFound(c, err)
	}
	return c.JSON(fiber.Map{
		"quiz_name": hosted.Quiz.Name,
		"username":  hosted.Quiz.User.Username,
	})
}

func (_playHandler) Play(c *fiber.Ctx) error {
	hosted := repository.HostedQuiz.FindByPublicKey(c.Params("public_key"))
	if hosted.Id == 0 {
		return errs.NotFound(c, errors.New("couldn't find this hosted quiz"))
	}

	questions := repository.Question.ForPlayers(hosted.QuizId)

	return c.JSON(questions)
}

func (_playHandler) Submit(c *fiber.Ctx) error {
	return c.SendString("Not implemented yet")
}

func (_playHandler) getHostedQuiz(publicKey string) (*models.HostedQuiz, error) {
	hosted := repository.HostedQuiz.FindByPublicKeyWithQuizUser(publicKey)
	if hosted.Id == 0 {
		return nil, errors.New("couldn't find this hosted quiz")
	}
	return hosted, nil
}
