package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type exampleController struct {
	exampleService model.ExampleService
}

func NewExampleController(exampleService model.ExampleService) model.ExampleController {
	return &exampleController{
		exampleService: exampleService,
	}
}

func (e *exampleController) HandleCreateExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateExampleRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := e.exampleService.CreateExample(c.Request().Context(), model.CreateExampleRequest{
			IntentID: req.IntentID,
			Example:  req.Example,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (e *exampleController) HandleFindAllExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		examples, err := e.exampleService.FindAllExampleByIntentID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, examples)
	}
}

func (e *exampleController) HandleFindExampleByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentId := c.Param("intentId")
		exampleId := c.Param("exampleId")

		example, err := e.exampleService.FindExampleByIntentID(c.Request().Context(), intentId, exampleId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, example)
	}
}

func (e *exampleController) HandleUpdateExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentId := c.Param("intentId")
		exampleId := c.Param("exampleId")

		req := model.UpdateExampleRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		example, err := e.exampleService.UpdateExample(c.Request().Context(), intentId, exampleId, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, example)
	}
}

func (e *exampleController) HandleDeleteExample() echo.HandlerFunc {
	return func(c echo.Context) error {
		intentId := c.Param("intentId")
		exampleId := c.Param("exampleId")

		isDeleted, err := e.exampleService.DeleteExample(c.Request().Context(), intentId, exampleId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
