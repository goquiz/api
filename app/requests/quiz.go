package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _quizValidation struct {
	Name string `json:"name" validate:"required,min=4,max=35,alphanumunicode"`
}

var QuizValidation _quizValidation

func QuizRequest(c *fiber.Ctx) error {
	r := &QuizValidation
	return Validate(r, c)
}
