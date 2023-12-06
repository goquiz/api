package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bndrmrtn/goquiz_api/app/repository"
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/database/models"
	"github.com/bndrmrtn/goquiz_api/helpers"
	"github.com/bndrmrtn/goquiz_api/http/errs"
	"github.com/bndrmrtn/goquiz_api/http/sessions"
	"github.com/gofiber/fiber/v2"
)

type _authHandler struct{}

var AuthHandler _authHandler

func (_authHandler) Login(c *fiber.Ctx) error {
	defer sessions.Global.Save()
	userRequest := requests.LoginValidation

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

	sessions.Global.Set("authorized.user_id", user.Id)
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("successfully logged in as %v", user.Username),
		"session": fiber.Map{
			"token": sessions.Global.Id(),
		},
	})
}

func (_authHandler) Register(c *fiber.Ctx) error {
	userRequest := requests.RegisterValidation
	defer sessions.Global.Save()

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
	sessions.Global.Set("authorized.user_id", user.Id)
	return c.JSON(fiber.Map{
		"message": "Successfully registered",
	})
}

func (_authHandler) Logout(c *fiber.Ctx) error {
	err := sessions.Global.Destroy()
	if err != nil {
		return errs.InternalServerError(c, err)
	}
	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
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
