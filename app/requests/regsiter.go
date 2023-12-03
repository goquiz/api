package requests

import (
	"github.com/gofiber/fiber/v2"
)

type register struct {
	Username string `json:"username" validate:"required,min=3,max=15"`
	Email    string `json:"email" validate:"required"`
}

var RegisterRequest register

func (r *register) Validator(c *fiber.Ctx) error {
	return Validate(r, c)
}
