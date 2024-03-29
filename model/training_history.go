package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type Status string

const (
	StatusDone   Status = "DONE"
	StatusFailed Status = "FAILED"
)

type CreateTrainingHistoryRequest struct {
	UserId    string `json:"user_id" validate:"required"`
	TotalTime int    `json:"total_time" validate:"required"`
	Status    Status `json:"status" validate:"required"`
}

func (c *CreateTrainingHistoryRequest) Validate() error {
	return validate.Struct(c)
}

type TrainingHistory struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	User      *User     `json:"user" gorm:"foreignKey:UserId"`
	TotalTime int       `json:"total_time"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TrainingHistoryController interface {
	HandleCreateTrainingHistory() echo.HandlerFunc
	HandleFindAllTrainingHistory() echo.HandlerFunc
}

type TrainingHistoryService interface {
	CreateTrainingHistory(ctx context.Context, req CreateTrainingHistoryRequest) (*TrainingHistory, error)
	FindAllTrainingHistory(ctx context.Context) ([]*TrainingHistory, error)
}

type TrainingHistoryRepository interface {
	Create(ctx context.Context, th *TrainingHistory) error
	FindAll(ctx context.Context) ([]*TrainingHistory, error)
}
