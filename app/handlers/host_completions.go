package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/http/errs"
	"strconv"
)

type _hostCompletionsHandler struct{}

var HostCompletionsHandler _hostCompletionsHandler

func (_hostCompletionsHandler) Paginate(c *fiber.Ctx) error {
	host, err := HostHandler.GetUserHost(c)

	if err != nil {
		return errs.BadRequest(c, err)
	}

	rawPage := c.Query("page")
	pageNum, _ := strconv.Atoi(rawPage)

	hosted := repository.HostedQuiz.PaginateSubmissions(
		host.Id,
		5,
		pageNum,
	)
	return c.JSON(hosted)
}

func (_hostCompletionsHandler) Analyze(*fiber.Ctx) error {
	return nil
}
