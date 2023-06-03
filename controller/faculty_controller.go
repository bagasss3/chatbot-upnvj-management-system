package controller

import (
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type facultyController struct {
	facultyService model.FacultyService
}

func NewFacultyController(facultyService model.FacultyService) model.FacultyController {
	return &facultyController{
		facultyService: facultyService,
	}
}

func (f *facultyController) HandleFindAllFaculty() echo.HandlerFunc {
	return func(c echo.Context) error {
		Facultys, err := f.facultyService.FindAllFaculty(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, Facultys)
	}
}

func (f *facultyController) HandleFindByIDFaculty() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		Faculty, err := f.facultyService.FindByIDFaculty(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, Faculty)
	}
}
