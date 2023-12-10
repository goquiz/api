package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/http/errs"
	"regexp"
	"strings"
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
			pattern := regexp.MustCompile("(\\p{Lu}+\\P{Lu}*)")
			s = pattern.ReplaceAllString(err.Field(), "${1}_")
			s, _ = strings.CutSuffix(strings.ToLower(err.Field()), "_")
			el.Field = s.(string)
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}
