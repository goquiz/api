package errs

import "github.com/gofiber/fiber/v2"

func BadRequest(c *fiber.Ctx, err error) error {
	status := fiber.StatusBadRequest
	return c.Status(status).JSON(map[string]interface{}{
		"error": map[string]interface{}{
			"code":    status,
			"message": err.Error(),
		},
	})
}
