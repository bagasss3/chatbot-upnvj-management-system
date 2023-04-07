package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

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
		req := model.CreateActionHttpRequest{}
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

func (a *actionHttpController) HandleFindActionHttpByActionID() echo.HandlerFunc {
	return func(c echo.Context) error {
		actionIdParam := c.Param("actionId")

		actionId, err := strconv.ParseInt(actionIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		actionHttp, err := a.actionHttpService.FindActionHttpByID(c.Request().Context(), actionId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, actionHttp)
	}
}

func (a *actionHttpController) HandleUpdateActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.UpdateActionHttpRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		actionIdParam := c.Param("actionId")
		actionId, err := strconv.ParseInt(actionIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		update, err := a.actionHttpService.UpdateActionHttp(c.Request().Context(), actionId, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (a *actionHttpController) HandleDeleteActionHttp() echo.HandlerFunc {
	return func(c echo.Context) error {
		actionIdParam := c.Param("actionId")

		actionId, err := strconv.ParseInt(actionIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := a.actionHttpService.DeleteActionHttp(c.Request().Context(), actionId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
