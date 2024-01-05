package requests

import "github.com/gofiber/fiber/v2"

type _playValidation struct {
	Answers []string `json:"answers" validate:"required,min=1,max=15,dive,min=1,max=35"`
}

var PlayValidation _playValidation

func PlayRequest(c *fiber.Ctx) error {
	r := &PlayValidation
	return Validate(r, c)
}
