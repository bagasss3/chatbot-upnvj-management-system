package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type entityController struct {
	entityService model.EntityService
}

func NewEntityController(entityService model.EntityService) model.EntityController {
	return &entityController{
		entityService: entityService,
	}
}

func (e *entityController) HandleCreateEntity() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateEntityRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := e.entityService.CreateEntity(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (e *entityController) HandleFindAllEntity() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentIdParam := c.Param("intentId")
		intentId, err := strconv.ParseInt(intentIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		entities, err := e.entityService.FindAllEntity(c.Request().Context(), intentId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, entities)
	}
}

func (e *entityController) HandleFindEntityByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentIdParam := c.Param("intentId")
		intentId, err := strconv.ParseInt(intentIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		entity, err := e.entityService.FindEntityByID(c.Request().Context(), id, intentId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, entity)
	}
}

func (e *entityController) HandleDeleteEntity() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentIdParam := c.Param("intentId")
		intentId, err := strconv.ParseInt(intentIdParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		entity, err := e.entityService.DeleteEntity(c.Request().Context(), id, intentId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, entity)
	}
}
