package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/http/authorized"
)

type _meHandler struct{}

var MeHandler _meHandler

// Hello returns the user information if logged in
func (_meHandler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("Hi %v!", authorized.Authorized.User.Username),
		"authUser": authorized.Authorized.User,
	})
}
