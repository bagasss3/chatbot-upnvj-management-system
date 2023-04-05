package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type actionController struct {
	actionService model.ActionService
}

func NewActionController(actionService model.ActionService) model.ActionController {
	return &actionController{
		actionService: actionService,
	}
}

func (a *actionController) HandleCreateAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateActionRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := a.actionService.CreateAction(c.Request().Context(), model.CreateUpdateActionRequest{
			Name: req.Name,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (a *actionController) HandleFindAllAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		actions, err := a.actionService.FindAllAction(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, actions)
	}
}

func (a *actionController) HandleFindActionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		action, err := a.actionService.FindActionByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, action)
	}
}

func (a *actionController) HandleUpdateAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateActionRequest{}
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

		update, err := a.actionService.UpdateAction(c.Request().Context(), id, model.CreateUpdateActionRequest{
			Name: req.Name,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (a *actionController) HandleDeleteAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := a.actionService.DeleteAction(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
