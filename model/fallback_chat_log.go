package model

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"

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
	Id        string         `json:"id" gorm:"column:id"`
	Chat      string         `json:"chat" gorm:"column:chat"`
	Cluster   null.Int       `json:"cluster" gorm:"column:cluster"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type ClusterData struct {
	Cluster null.Int           `json:"cluster"`
	Data    []*FallbackChatLog `json:"data" gorm:"-"`
}

type UpdateGroupClusterRequest struct {
	Cluster int `json:"cluster" validate:"required"`
}

type ResponseFallback struct {
	ExistingLog []*FallbackChatLog `json:"existing_log"`
	NewLog      []*FallbackChatLog `json:"new_log"`
}

func (c *UpdateGroupClusterRequest) Validate() error {
	return validate.Struct(c)
}

type FallbackChatLogController interface {
	HandleCreateFallbackChatLog() echo.HandlerFunc
	HandleFindAllFallbackChatLog() echo.HandlerFunc
	HandleFindAllFallbackChatLogGroupCluster() echo.HandlerFunc
	HandleFindAllFallbackChatLogOldAndNew() echo.HandlerFunc
	HandleUpdateGroupCluster() echo.HandlerFunc
}

type FallbackChatLogService interface {
	CreateFallbackChatLog(ctx context.Context, req CreateFallbackChatLogRequest) (*FallbackChatLog, error)
	FindAllFallbackChatLog(ctx context.Context, page string) ([]*FallbackChatLog, error)
	FindAllFallbackChatLogGroupCluster(ctx context.Context) ([]*ClusterData, error)
	FindAllFallbackChatLogOldAndNew(ctx context.Context) (*ResponseFallback, error)
	UpdateGroupCluster(ctx context.Context) ([]*ClusterData, error)
}

type FallbackChatLogRepository interface {
	Create(ctx context.Context, fcl *FallbackChatLog) error
	FindAll(ctx context.Context, page string) ([]*FallbackChatLog, error)
	FindAllGroupCluster(ctx context.Context) ([]*ClusterData, error)
	FindNullAndCluster(ctx context.Context) (*ResponseFallback, error)
	UpdateGroupCluster(ctx context.Context, fallback *FallbackChatLog, tx *gorm.DB) error
}
