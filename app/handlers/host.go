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

type _host struct{}

var HostHandler _host

func (_host) New(c *fiber.Ctx) error {
	quiz, err := QuizHandler.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	hostId := repository.HostedQuiz.NewUniqueCode()
	hostedQuiz := models.HostedQuiz{
		PublicKey: hostId,
		QuizId:    quiz.Id,
		IsActive:  true,
		Name:      requests.HostValidation.Name,
	}
	database.Database.Create(&hostedQuiz)

	if hostedQuiz.Id == 0 {
		return errs.InternalServerError(c, errors.New("failed to host this quiz"))
	}

	return c.JSON(fiber.Map{
		"message": "Successfully hosted this quiz",
		"host":    hostedQuiz,
	})
}

func (_host) All(c *fiber.Ctx) error {
	idInt, _ := strconv.Atoi(c.Params("id"))
	quizId := uint(idInt)
	userId := authorized.Authorized.User.Id
	hostedQuizzes := repository.HostedQuiz.AllForUser(quizId, userId)
	return c.JSON(fiber.Map{
		"hosts": hostedQuizzes,
		"quiz":  repository.Quiz.ById(quizId),
	})
}

func (h _host) ChangeActive(c *fiber.Ctx) error {
	type IsActive struct {
		IsActive bool `json:"is_active"`
	}
	var isActive IsActive

	err := c.BodyParser(&isActive)
	if err != nil {
		return errs.BadRequest(c, err)
	}
	hosted, err := h.GetUserHost(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	hosted.IsActive = isActive.IsActive

	if !isActive.IsActive {
		hosted.PublicKey = ""
	} else {
		hosted.PublicKey = repository.HostedQuiz.NewUniqueCode()
	}

	database.Database.Save(&hosted)
	return c.JSON(fiber.Map{
		"message":    "Successfully modified",
		"is_active":  hosted.IsActive,
		"public_key": hosted.PublicKey,
	})
}

func (h _host) Destroy(c *fiber.Ctx) error {
	hosted, err := h.GetUserHost(c)
	if err != nil {
		return errs.NotFound(c, err)
	}
	database.Database.Delete(&hosted)
	return c.JSON(fiber.Map{
		"message": "successfully deleted",
	})
}

func (_host) GetUserHost(c *fiber.Ctx) (*models.HostedQuiz, error) {
	idInt, _ := strconv.Atoi(c.Params("hostId"))
	id := uint(idInt)
	hosted := repository.HostedQuiz.FindForUser(id, authorized.Authorized.User.Id)
	if hosted.Id == 0 {
		return nil, errors.New("this host could not be found or does not belongs to you")
	}
	return hosted, nil
}
