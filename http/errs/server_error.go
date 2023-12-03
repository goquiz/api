package errs

import "github.com/gofiber/fiber/v2"

func InternalServerError(c *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	return c.Status(status).JSON(map[string]interface{}{
		"error": map[string]interface{}{
			"code":    status,
			"message": err.Error(),
		},
	})
}
