package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type intentController struct {
	intentService model.IntentService
}

func NewIntentController(intentService model.IntentService) model.IntentController {
	return &intentController{
		intentService: intentService,
	}
}

func (i *intentController) HandleCreateIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateIntentRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := i.intentService.CreateIntent(c.Request().Context(), model.CreateUpdateIntentRequest{
			Name: req.Name,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (i *intentController) HandleFindAllIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		intents, err := i.intentService.FindAllIntent(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, intents)
	}
}

func (i *intentController) HandleFindIntentByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		intent, err := i.intentService.FindIntentByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, intent)
	}
}

func (i *intentController) HandleUpdateIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateIntentRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		update, err := i.intentService.UpdateIntent(c.Request().Context(), id, model.CreateUpdateIntentRequest{
			Name: req.Name,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (i *intentController) HandleDeleteIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := i.intentService.DeleteIntent(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}

func (i *intentController) HandleCountAllIntent() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := i.intentService.CountAllIntent(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, count)
	}
}
