package handlers

import (
	"github.com/bndrmrtn/goquiz_api/http/authorized"
	"github.com/gofiber/fiber/v2"
)

type me struct{}

var Me me

func (me) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello " + authorized.Authorized.User.Username,
	})
}
