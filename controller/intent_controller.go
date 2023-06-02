package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

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
		name := c.QueryParam("name")
		intents, err := i.intentService.FindAllIntent(c.Request().Context(), name)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, intents)
	}
}

func (i *intentController) HandleFindAllWithExamples() echo.HandlerFunc {
	return func(c echo.Context) error {
		intents, err := i.intentService.FindAllWithExamples(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, intents)
	}
}

func (i *intentController) HandleFindIntentByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

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

		id := c.Param("id")

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
		id := c.Param("id")

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
