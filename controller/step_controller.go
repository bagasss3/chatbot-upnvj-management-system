package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type stepController struct {
	stepService model.StepService
}

func NewStepController(stepService model.StepService) model.StepController {
	return &stepController{
		stepService: stepService,
	}
}

func (s *stepController) HandleCreateStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateStepArrayRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := s.stepService.CreateStep(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (s *stepController) HandleFindAllStepByStoryID() echo.HandlerFunc {
	return func(c echo.Context) error {
		storyId := c.Param("storyId")

		stories, err := s.stepService.FindAllStepByStoryID(c.Request().Context(), storyId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, stories)
	}
}

func (s *stepController) HandleFindStepByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		storyId := c.Param("storyId")

		step, err := s.stepService.FindStepByID(c.Request().Context(), id, storyId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, step)
	}
}

func (s *stepController) HandleUpdateStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.UpdateStepRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		id := c.Param("id")
		storyId := c.Param("storyId")

		update, err := s.stepService.UpdateStep(c.Request().Context(), id, storyId, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (s *stepController) HandleDeleteStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		storyId := c.Param("storyId")

		isDeleted, err := s.stepService.DeleteStep(c.Request().Context(), id, storyId)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
