package handlers

import (
	"github.com/bndrmrtn/goquiz_api/app/requests"
	"github.com/gofiber/fiber/v2"
)

type auth struct{}

var Auth auth

func (auth) Login(*fiber.Ctx) error {
	return nil
}

func (auth) Register(c *fiber.Ctx) error {
	body := requests.RegisterRequest
	c.BodyParser(&body)
	return c.JSON(body)
}
