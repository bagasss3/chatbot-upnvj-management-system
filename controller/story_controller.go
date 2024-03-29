package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type storyController struct {
	storyService model.StoryService
}

func NewStoryController(storyService model.StoryService) model.StoryController {
	return &storyController{
		storyService: storyService,
	}
}

func (s *storyController) HandleCreateStory() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateStoryRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := s.storyService.CreateStory(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (s *storyController) HandleFindAllStory() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		stories, err := s.storyService.FindAllStory(c.Request().Context(), name)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, stories)
	}
}

func (s *storyController) HandleFindStoryByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		story, err := s.storyService.FindStoryByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, story)
	}
}

func (s *storyController) HandleUpdateStory() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateUpdateStoryRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		id := c.Param("id")

		update, err := s.storyService.UpdateStory(c.Request().Context(), id, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (s *storyController) HandleDeleteStory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		isDeleted, err := s.storyService.DeleteStory(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}
