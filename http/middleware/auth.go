package middleware

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/http/errs"
	"github.com/goquiz/api/http/sessions"
)

type _authMiddleware struct{}

var AuthMiddleware _authMiddleware

func (_authMiddleware) Auth(c *fiber.Ctx) error {
	if sessions.Global.Session == nil {
		return errs.Unauthorized(c, errors.New("failed to authenticate through session"))
	}
	userId := sessions.Global.Get("authorized.user_id")
	if userId != nil && fmt.Sprintf("%T", userId) == "uint" {
		user, err := repository.User.FindById(userId.(uint))
		if err != nil {
			return errs.Unauthorized(c, errors.New("failed to authenticate through this session"))
		}
		c.Locals("auth.user", user)
		return c.Next()
	}
	return errs.Unauthorized(c, errors.New("failed to authenticate through this session"))
}
