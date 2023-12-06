package requests

import (
	"github.com/bndrmrtn/goquiz_api/http/errs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string `json:"field,omitempty"`
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

var Validator = validator.New()

func Validate(s interface{}, c *fiber.Ctx) error {
	var errors []*IError
	err := c.BodyParser(&s)
	if err != nil {
		return errs.InternalServerError(c, err)
	}
	err = Validator.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
