package controller

import (
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type workerController struct {
	workerService model.WorkerService
}

func NewWorkerController(workerService model.WorkerService) model.WorkerController {
	return &workerController{
		workerService: workerService,
	}
}

func (w *workerController) HandleStartTrainingModel() echo.HandlerFunc {
	return func(c echo.Context) error {
		create, err := w.workerService.StartTrainingModel(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}
