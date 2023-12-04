package middleware

import (
	"errors"
	"fmt"
	"github.com/bndrmrtn/goquiz_api/http/authorized"
	"github.com/bndrmrtn/goquiz_api/http/errs"
	"github.com/bndrmrtn/goquiz_api/http/sessions"
	"github.com/gofiber/fiber/v2"
)

type _authMiddleware struct{}

var AuthMiddleware _authMiddleware

func (_authMiddleware) Auth(c *fiber.Ctx) error {
	if sessions.Global.Session == nil {
		return errs.Unauthorized(c, errors.New("failed to authenticate through session"))
	}
	userId := sessions.Global.Get("authorized.user_id")
	if userId != nil && fmt.Sprintf("%T", userId) == "uint" {
		err := authorized.Authorized.AuthUser(userId.(uint))
		if err != nil {
			return errs.Unauthorized(c, errors.New("failed to authenticate through this session"))
		}
		return c.Next()
	}
	return errs.Unauthorized(c, errors.New("failed to authenticate through this session"))
}
