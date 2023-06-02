package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type actionHttpController struct {
	actionHttpService model.ActionHttpService
}

func NewActionHttpController(actionHttpService model.ActionHttpService) model.ActionHttpController {
	return &actionHttpController{
		actionHttpService: actionHttpService,
	}
}

func (a *actionHttpController) HandleCreateActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateActionHttpRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := a.actionHttpService.CreateActionHttp(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (a *actionHttpController) HandleFindAllActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		actionHttps, err := a.actionHttpService.FindAllActionHttp(c.Request().Context(), name)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, actionHttps)
	}
}

func (a *actionHttpController) HandleFindAllActionHttpWithReqBodies() echo.HandlerFunc {
	return func(c echo.Context) error {

		actionHttps, err := a.actionHttpService.FindAllWithReqBodies(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, actionHttps)
	}
}

func (a *actionHttpController) HandleFindActionHttpByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		actionHttp, err := a.actionHttpService.FindActionHttpByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, actionHttp)
	}
}

func (a *actionHttpController) HandleUpdateActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateActionHttpRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		id := c.Param("id")

		update, err := a.actionHttpService.UpdateActionHttp(c.Request().Context(), id, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (a *actionHttpController) HandleDeleteActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		isDeleted, err := a.actionHttpService.DeleteActionHttp(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}

func (a *actionHttpController) HandleCountAllActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := a.actionHttpService.CountAllActionHttp(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, count)
	}
}
