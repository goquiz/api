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

type _host struct{}

var HostHandler _host

// New creates a new host for a quiz
func (_host) New(c *fiber.Ctx) error {
	quiz, err := QuizHandler.GetQuiz(c)
	if err != nil {
		return errs.NotFound(c, err)
	}

	quizHostsCount := repository.HostedQuiz.CountForQuizId(quiz.Id)
	if quizHostsCount >= 5 {
		return errs.BadRequest(c, errors.New("cannot have more than 5 hosts for a quiz"))
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

// All returns all the hosts for a given quiz
func (_host) All(c *fiber.Ctx) error {
	idInt, _ := strconv.Atoi(c.Params("id"))
	quizId := uint(idInt)
	userId := GetAuthUser(c).Id
	quiz := repository.Quiz.ById(quizId)

	if quiz.Id == 0 {
		return errs.NotFound(c, errors.New("couldn't find this quiz"))
	}

	hostedQuizzes := repository.HostedQuiz.AllForUser(quizId, userId)
	return c.JSON(fiber.Map{
		"hosts": hostedQuizzes,
		"quiz":  quiz,
	})
}

// ChangeActive changes the quiz activity (it sets or deletes the publicKey)
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

// Destroy deletes a hosted_quiz
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

// GetUserHost returns a models.HostedQuiz by the route "hostId" parameter
func (_host) GetUserHost(c *fiber.Ctx) (*models.HostedQuiz, error) {
	idInt, _ := strconv.Atoi(c.Params("hostId"))
	id := uint(idInt)
	hosted := repository.HostedQuiz.FindForUser(id, GetAuthUser(c).Id)
	if hosted.Id == 0 {
		return nil, errors.New("this host could not be found or does not belongs to you")
	}
	return hosted, nil
}
