package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _hostValidation struct {
	Name string `json:"name" validate:"required,min=4,max=35"`
}

var HostValidation _hostValidation

func HostRequest(c *fiber.Ctx) error {
	r := &HostValidation
	return Validate(r, c)
}
