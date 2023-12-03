package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/bndrmrtn/goquiz_api/app/repository"
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/database/models"
	"github.com/bndrmrtn/goquiz_api/helpers"
	"github.com/bndrmrtn/goquiz_api/http/errs"
	"github.com/gofiber/fiber/v2"
)

type auth struct{}

var Auth auth

func (auth) Login(c *fiber.Ctx) error {
	userRequest := requests.LoginValidation{}
	_ = c.BodyParser(&userRequest)
	user, err := repository.User.FindByUsername(userRequest.Username)
	if err != nil {
		return errs.Unauthorized(c, err)
	}

	password, _, err := passwordHash(userRequest.Password, user.PasswordSalt)
	if err != nil {
		return errs.InternalServerError(c, err)
	}

	if password != user.Password {
		return errs.Unauthorized(c, errors.New("invalid username or password"))
	}

	return c.SendString("successfully logged in as " + user.Username)
}

func (auth) Register(c *fiber.Ctx) error {
	userRequest := requests.RegisterValidation{}
	_ = c.BodyParser(&userRequest)

	if repository.User.IsUsernameOrEmailExists(userRequest.Username, userRequest.Email) {
		return errs.BadRequest(c, errors.New("user already exists with this username or email address"))
	}

	password, passwordSalt, err := passwordHash(userRequest.Password, "")

	if err != nil {
		return errs.InternalServerError(c, err)
	}

	user := models.User{
		Username:     userRequest.Username,
		Email:        userRequest.Email,
		Password:     password,
		PasswordSalt: passwordSalt,
	}

	database.Database.Model(&models.User{}).Create(&user)
	return c.SendString("Successfully registered")
}

func passwordHash(p string, s string) (string, string, error) {
	if s == "" {
		s = helpers.NewRandom().String(25)
	}

	hash := sha256.New()
	var err error
	_, err = hash.Write([]byte(p))
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hash.Sum([]byte(s))), s, nil
}
