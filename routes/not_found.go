package routes

import (
	"errors"
	"github.com/bndrmrtn/goquiz_api/http/errs"
	"github.com/gofiber/fiber/v2"
)

type _notFound struct{}

var NotFoundPage _notFound

func (_notFound) New(c *fiber.Ctx) error {
	return errs.NotFound(c, errors.New("this page could not be found"))
}
