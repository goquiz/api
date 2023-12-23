package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _resetPasswordValidation struct {
	Username string `json:"username" validate:"required,min=3,max=15,alphanumunicode"`
}

var ResetPasswordValidation _resetPasswordValidation

func ResetPasswordRequest(c *fiber.Ctx) error {
	r := &ResetPasswordValidation
	return Validate(r, c)
}
