package requests

import (
	"github.com/gofiber/fiber/v2"
)

type _questionValidation struct {
	Question string   `json:"question" validate:"required,min=5,max=55"`
	Image    *string  `json:"image,omitempty"`
	Answers  []string `json:"answers" validate:"required,min=1,max=4,dive,min=1,max=35"`
	Answer   string   `json:"answer" validate:"required,min=1,max=35"`
}

var QuestionValidation _questionValidation

func QuestionRequest(c *fiber.Ctx) error {
	r := &QuestionValidation
	return Validate(r, c)
}
