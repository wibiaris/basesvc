package author

import (
	"net/http"

	"bitbucket.com/kudoindonesia/imgart/log"
	"github.com/RoseRocket/xerrs"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/internal/respond"
)

// GetAll handles HTTP GET request for retrieving multiple authors
func GetAll(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()

		res, err := appController.Author.GetAuthors(ctx.Context())
		if err != nil {
			logger.LogEntrySetFields(ctx, log.Fields{
				"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
				"context":     "GetAuthors",
				"resp_status": http.StatusInternalServerError,
			})
			return respond.Fail(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err)

		}

		return respond.Success(ctx, http.StatusOK, res)
	}
}
