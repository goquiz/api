package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/http/errs"
)

func NewNotFoundPage(c *fiber.Ctx) error {
	return errs.NotFound(c, errors.New("this page could not be found"))
}
