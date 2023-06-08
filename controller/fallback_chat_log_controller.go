package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type fallbackChatLogController struct {
	fallbackChatLogService model.FallbackChatLogService
}

func NewFallbackChatLogController(fallbackChatLogService model.FallbackChatLogService) model.FallbackChatLogController {
	return &fallbackChatLogController{
		fallbackChatLogService: fallbackChatLogService,
	}
}

func (fcl *fallbackChatLogController) HandleCreateFallbackChatLog() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateFallbackChatLogRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := fcl.fallbackChatLogService.CreateFallbackChatLog(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (fcl *fallbackChatLogController) HandleFindAllFallbackChatLog() echo.HandlerFunc {
	return func(c echo.Context) error {
		page := c.QueryParam("page")
		fallbacks, err := fcl.fallbackChatLogService.FindAllFallbackChatLog(c.Request().Context(), page)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, fallbacks)
	}
}

func (fcl *fallbackChatLogController) HandleFindAllFallbackChatLogGroupCluster() echo.HandlerFunc {
	return func(c echo.Context) error {
		fallbacks, err := fcl.fallbackChatLogService.FindAllFallbackChatLogGroupCluster(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, fallbacks)
	}
}

func (fcl *fallbackChatLogController) HandleFindAllFallbackChatLogOldAndNew() echo.HandlerFunc {
	return func(c echo.Context) error {
		fallbacks, err := fcl.fallbackChatLogService.FindAllFallbackChatLogOldAndNew(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, fallbacks)
	}
}

func (fcl *fallbackChatLogController) HandleUpdateGroupCluster() echo.HandlerFunc {
	return func(c echo.Context) error {
		fallbacks, err := fcl.fallbackChatLogService.UpdateGroupCluster(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, fallbacks)
	}
}
