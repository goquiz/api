package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type _meHandler struct{}

var MeHandler _meHandler

// Hello returns the user information if logged in
func (_meHandler) Hello(c *fiber.Ctx) error {
	authUser := GetAuthUser(c)
	return c.JSON(fiber.Map{
		"message":  fmt.Sprintf("Hi %v!", authUser.Username),
		"authUser": authUser,
	})
}
