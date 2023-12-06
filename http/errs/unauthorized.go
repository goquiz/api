package errs

import "github.com/gofiber/fiber/v2"

func Unauthorized(c *fiber.Ctx, err error) error {
	status := fiber.StatusUnauthorized
	return c.Status(status).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    status,
			"message": err.Error(),
		},
	})
}
