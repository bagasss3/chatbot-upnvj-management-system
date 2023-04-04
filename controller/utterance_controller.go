package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type utteranceController struct {
	utteranceService model.UtteranceService
}

func NewUtteranceController(utteranceService model.UtteranceService) model.UtteranceController {
	return &utteranceController{
		utteranceService: utteranceService,
	}
}

func (u *utteranceController) HandleCreateUtterance() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateUtteranceRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := u.utteranceService.CreateUtterance(c.Request().Context(), model.CreateUpdateUtteranceRequest{
			Name:     req.Name,
			Response: req.Response,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (u *utteranceController) HandleFindAllUtterance() echo.HandlerFunc {
	return func(c echo.Context) error {
		utterances, err := u.utteranceService.FindAllUtterance(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, utterances)
	}
}

func (u *utteranceController) HandleFindUtteranceByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		utterance, err := u.utteranceService.FindUtteranceByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, utterance)
	}
}

func (u *utteranceController) HandleUpdateUtterance() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateUtteranceRequest{}
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

		update, err := u.utteranceService.UpdateUtterance(c.Request().Context(), id, model.CreateUpdateUtteranceRequest{
			Name:     req.Name,
			Response: req.Response,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (u *utteranceController) HandleDeleteUtterance() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := u.utteranceService.DeleteUtterance(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
