package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type configurationController struct {
	configurationService model.ConfigurationService
}

func NewConfigurationController(configurationService model.ConfigurationService) model.ConfigurationController {
	return &configurationController{
		configurationService: configurationService,
	}
}

func (co *configurationController) HandleCreateConfiguration() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateConfigurationRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := co.configurationService.CreateConfiguration(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (co *configurationController) HandleFindConfiguration() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		conf, err := co.configurationService.FindConfiguration(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, conf)
	}
}

func (co *configurationController) HandleUpdateConfiguration() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateConfigurationRequest{}
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

		update, err := co.configurationService.UpdateConfiguration(c.Request().Context(), id, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}
