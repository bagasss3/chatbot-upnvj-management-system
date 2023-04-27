package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type krsActionController struct {
	krsActionService model.KrsActionService
}

func NewKrsActionController(krsActionService model.KrsActionService) model.KrsActionController {
	return &krsActionController{
		krsActionService: krsActionService,
	}
}

func (k *krsActionController) HandleCreateKrsAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateKrsActionRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := k.krsActionService.CreateKrsAction(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (k *krsActionController) HandleFindAllKrsAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		krsActions, err := k.krsActionService.FindAllKrsAction(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, krsActions)
	}
}

func (k *krsActionController) HandleFindKrsActionByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		krsAction, err := k.krsActionService.FindKrsActionByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, krsAction)
	}
}

func (k *krsActionController) HandleUpdateKrsAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateKrsActionRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		id := c.Param("id")

		update, err := k.krsActionService.UpdateKrsAction(c.Request().Context(), id, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (k *krsActionController) HandleDeleteKrsAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		isDeleted, err := k.krsActionService.DeleteKrsAction(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
