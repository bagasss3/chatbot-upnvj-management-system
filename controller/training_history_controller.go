package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type trainingHistoryController struct {
	trainingHistoryService model.TrainingHistoryService
}

func NewTrainingHistoryController(trainingHistoryService model.TrainingHistoryService) model.TrainingHistoryController {
	return &trainingHistoryController{
		trainingHistoryService: trainingHistoryService,
	}
}

func (th *trainingHistoryController) HandleCreateTrainingHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateTrainingHistoryRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := th.trainingHistoryService.CreateTrainingHistory(c.Request().Context(), req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}

func (th *trainingHistoryController) HandleFindAllTrainingHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		trainingHistories, err := th.trainingHistoryService.FindAllTrainingHistory(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, trainingHistories)
	}
}
