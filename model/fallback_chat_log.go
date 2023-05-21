package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FallbackPage string

const (
	FallbackPageDashboard FallbackPage = "DASHBOARD"
	FallbackPageLog       FallbackPage = "LOG"
)

type CreateFallbackChatLogRequest struct {
	Chat string `json:"chat" validate:"required"`
}

func (c *CreateFallbackChatLogRequest) Validate() error {
	return validate.Struct(c)
}

type FallbackChatLog struct {
	Id        string         `json:"id"`
	Chat      string         `json:"chat"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type FallbackChatLogController interface {
	HandleCreateFallbackChatLog() echo.HandlerFunc
	HandleFindAllFallbackChatLog() echo.HandlerFunc
}

type FallbackChatLogService interface {
	CreateFallbackChatLog(ctx context.Context, req CreateFallbackChatLogRequest) (*FallbackChatLog, error)
	FindAllFallbackChatLog(ctx context.Context, page string) ([]*FallbackChatLog, error)
}

type FallbackChatLogRepository interface {
	Create(ctx context.Context, fcl *FallbackChatLog) error
	FindAll(ctx context.Context, page string) ([]*FallbackChatLog, error)
}
