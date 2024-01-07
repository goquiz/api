package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _resetPasswordValidation struct {
	Password string `json:"password" validate:"required,min=10,max=55"`
}

var ResetPasswordValidation _resetPasswordValidation

func ResetPasswordRequest(c *fiber.Ctx) error {
	r := &ResetPasswordValidation
	return Validate(r, c)
}
