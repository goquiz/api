package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _requestNewPasswordValidation struct {
	Username string `json:"username" validate:"required,min=3,max=15,alphanumunicode"`
}

var RequestNewPasswordValidation _requestNewPasswordValidation

func RequestNewPasswordRequest(c *fiber.Ctx) error {
	r := &RequestNewPasswordValidation
	return Validate(r, c)
}
