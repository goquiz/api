package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _loginValidation struct {
	Username string `json:"username" validate:"required,min=3,max=15,alphanumunicode"`
	Password string `json:"password" validate:"required,min=10,max=25"`
}

var LoginValidation _loginValidation

func LoginRequest(c *fiber.Ctx) error {
	r := &LoginValidation
	return Validate(r, c)
}
