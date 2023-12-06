package errs

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx, err error) error {
	status := fiber.StatusNotFound
	return c.Status(status).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    status,
			"message": err.Error(),
		},
	})
}
