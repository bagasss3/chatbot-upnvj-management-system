package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type logIntentController struct {
	logIntentService model.LogIntentService
}

func NewLogIntentController(logIntentService model.LogIntentService) model.LogIntentController {
	return &logIntentController{
		logIntentService: logIntentService,
	}
}

func (li *logIntentController) HandleCreateOrUpdateLogIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateLogIntentRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := li.logIntentService.CreateOrUpdateLogIntent(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (li *logIntentController) HandleFindAllLogIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		logIntents, err := li.logIntentService.FindAllLogIntent(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, logIntents)
	}
}

func (li *logIntentController) HandleFindLogIntentByIntentID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("intentId")

		logIntent, err := li.logIntentService.FindLogIntentByIntentID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, logIntent)
	}
}
