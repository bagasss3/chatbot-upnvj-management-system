package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type WorkerController interface {
	HandleStartTrainingModel() echo.HandlerFunc
}

type WorkerService interface {
	StartTrainingModel(ctx context.Context) (*TrainingHistory, error)
}
