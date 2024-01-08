package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/app/requests"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
	"github.com/goquiz/api/http/errs"
	"github.com/goquiz/api/http/sessions"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

type _authHandler struct{}

var AuthHandler _authHandler

// Login authorize the user's credentials and returns a session if successful
func (_authHandler) Login(c *fiber.Ctx) error {
	defer sessions.Global.Save()
	userRequest := requests.LoginValidation

	user, err := repository.User.FindByUsername(strings.ToLower(userRequest.Username))
	if err != nil {
		return errs.Unauthorized(c, err)
	}

	if !passwordCompare(user.Password, userRequest.Password, user.PasswordSalt) {
		return errs.Unauthorized(c, errors.New("invalid username or password"))
	}

	if user.EmailVerifiedAt == nil {
		return errs.BadRequest(c, errors.New("verify your email address before login"))
	}

	sessions.Global.Set("authorized.user_id", user.Id)
	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("successfully logged in as %v", user.Username),
		"authUser": user,
		"session":  sessions.Global.Id(),
	})
}

// Register creates a new user with the given credentials
func (_authHandler) Register(c *fiber.Ctx) error {
	userRequest := requests.RegisterValidation
	defer sessions.Global.Save()

	if repository.User.IsUsernameOrEmailExists(strings.ToLower(userRequest.Username), strings.ToLower(userRequest.Email)) {
		return errs.BadRequest(c, errors.New("user already exists with this username or email address"))
	}

	password, passwordSalt, err := passwordHash(userRequest.Password, "")

	if err != nil {
		return errs.InternalServerError(c, err)
	}

	user := models.User{
		Username:     strings.ToLower(userRequest.Username),
		Email:        userRequest.Email,
		Password:     password,
		PasswordSalt: passwordSalt,
	}

	database.Database.Model(&models.User{}).Create(&user)

	emailVerificationToken := repository.User.NewTokenFor(&models.EmailVerification{})

	emailVerification := models.EmailVerification{
		UserId:     user.Id,
		Token:      emailVerificationToken,
		Expiration: time.Now().Add(time.Hour * 3),
	}

	database.Database.Save(&emailVerification)

	mail := helpers.NewMail("Quizzes.LOL<noreply@quizzes.lol>", user.Email)
	mail.Subject("Email verification")
	mail.Body(
		fmt.Sprintf("Hi %v,<br/>Please click <a href=\"https://quizzes.lol/email-verification/%v\">here</a> to verify your email address.<br/><br/>- We never ask for passwords or credentials via email.", cases.Title(language.English).String(user.Username), emailVerificationToken),
		true,
	)

	err = mail.Send()
	if err != nil {
		return errs.InternalServerError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Successfully registered",
	})
}

// Logout logs out the user
func (_authHandler) Logout(c *fiber.Ctx) error {
	err := sessions.Global.Destroy()
	if err != nil {
		return errs.InternalServerError(c, err)
	}
	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}

// VerifyEmailAddress verifies the email for a user to be able to use the app
func (_authHandler) VerifyEmailAddress(c *fiber.Ctx) error {
	token := c.Params("token")
	emailVerification, err := repository.User.EmailVerification(token)
	if err != nil {
		return errs.NotFound(c, err)
	}

	user := emailVerification.User
	verifiedAt := time.Now()
	user.EmailVerifiedAt = &verifiedAt
	database.Database.Save(&user)
	database.Database.Delete(&emailVerification)

	return c.JSON(fiber.Map{
		"message": "Successfully verified",
	})
}

func (_authHandler) RequestNewPassword(c *fiber.Ctx) error {
	username := strings.ToLower(requests.RequestNewPasswordValidation.Username)
	user, err := repository.User.FindByUsername(username)
	if err != nil {
		return errs.NotFound(c, err)
	}

	if repository.User.HasRequestedNewPassword(user.Id) == true {
		return errs.BadRequest(c, errors.New("you have already requested a password reset"))
	}

	resetPasswordToken := repository.User.NewTokenFor(&models.ResetPassword{})

	resetPassword := &models.ResetPassword{
		UserId:     user.Id,
		Token:      resetPasswordToken,
		Expiration: time.Now().Add(time.Hour * 3),
	}
	database.Database.Create(&resetPassword)

	mail := helpers.NewMail("Quizzes.LOL<noreply@quizzes.lol>", user.Email)
	mail.Subject("Reset your password")
	mail.Body(
		fmt.Sprintf("Hi %v,<br/>Please click <a href=\"https://quizzes.lol/reset-password/%v\">here</a> to change your password.<br/><br/>- We never ask for passwords or credentials via email.", cases.Title(language.English).String(user.Username), resetPasswordToken),
		true,
	)

	err = mail.Send()
	if err != nil {
		return errs.InternalServerError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Reset password email sent",
	})
}

func (_authHandler) ChangePassword(c *fiber.Ctx) error {
	token := c.Params("token")
	resetPassword, err := repository.User.ResetPassword(token)

	if err != nil {
		return errs.NotFound(c, err)
	}

	rawPassword := requests.ResetPasswordValidation.Password
	user := resetPassword.User
	password, salt, err := passwordHash(rawPassword, "")

	if err != nil {
		return errs.InternalServerError(c, err)
	}

	user.Password = password
	user.PasswordSalt = salt
	database.Database.Save(&user)

	return c.JSON(fiber.Map{
		"message": "Password successfully changed",
	})
}

// passwordHash hashes the user's password with a given, or a random salt
func passwordHash(p string, s string) (string, string, error) {
	if s == "" {
		s = helpers.NewRandom().String(25)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(p+s), 14)

	if err != nil {
		return "", "", err
	}

	return string(bytes), s, nil
}

func passwordCompare(h string, p string, s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p+s))
	return err == nil
}
