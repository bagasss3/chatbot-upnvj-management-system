package controller

import (
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type majorController struct {
	majorService model.MajorService
}

func NewMajorController(majorService model.MajorService) model.MajorController {
	return &majorController{
		majorService: majorService,
	}
}

func (m *majorController) HandleFindAllMajor() echo.HandlerFunc {
	return func(c echo.Context) error {
		majors, err := m.majorService.FindAllMajor(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, majors)
	}
}

func (m *majorController) HandleFindByIDMajor() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		major, err := m.majorService.FindByIDMajor(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, major)
	}
}
