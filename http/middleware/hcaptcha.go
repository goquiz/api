package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/helpers"
	"github.com/goquiz/api/http/errs"
	"net/http"
	"net/url"
)

const (
	HCaptchaUrl = "https://api.hcaptcha.com/siteverify"
)

type hCaptcha struct{}

var HCaptcha hCaptcha

func (h hCaptcha) New(c *fiber.Ctx) error {
	data := struct {
		HCaptchaToken string `json:"hcaptcha_token"`
	}{}

	m := errors.New("unable to check that you are not a robot")

	err := c.BodyParser(&data)
	if err != nil {
		return errs.BadRequest(c, m)
	}

	status, err := h.Validate(data.HCaptchaToken)
	if err != nil {
		return errs.BadRequest(c, err)
	}

	if status {
		return c.Next()
	} else {
		return errs.BadRequest(c, m)
	}
}

func (hCaptcha) Validate(token string) (bool, error) {
	res, err := http.PostForm(HCaptchaUrl, url.Values{
		"secret":   {helpers.Env.HCaptcha.SecretKey},
		"response": {token},
	})

	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	o := struct {
		Success bool `json:"success"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&o)

	if err != nil {
		return false, err
	}

	return o.Success, nil
}
