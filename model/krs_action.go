package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateKrsActionRequest struct {
	Name       string `json:"name" validate:"required,min=3,max=60"`
	GetHttpReq string `json:"get_http_req" validate:"required"`
	ApiKey     string `json:"api_key" validate:"omitempty"`
}

func (c *CreateUpdateKrsActionRequest) Validate() error {
	return validate.Struct(c)
}

type KrsAction struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	GetHttpReq string         `json:"get_http_req"`
	ApiKey     string         `json:"api_key"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type KrsActionController interface {
	HandleCreateKrsAction() echo.HandlerFunc
	HandleFindAllKrsAction() echo.HandlerFunc
	HandleFindKrsActionByID() echo.HandlerFunc
	HandleUpdateKrsAction() echo.HandlerFunc
	HandleDeleteKrsAction() echo.HandlerFunc
}

type KrsActionService interface {
	CreateKrsAction(ctx context.Context, req CreateUpdateKrsActionRequest) (*KrsAction, error)
	FindAllKrsAction(ctx context.Context) ([]*KrsAction, error)
	FindKrsActionByID(ctx context.Context, id string) (*KrsAction, error)
	UpdateKrsAction(ctx context.Context, id string, req CreateUpdateKrsActionRequest) (*KrsAction, error)
	DeleteKrsAction(ctx context.Context, id string) (bool, error)
}

type KrsActionRepository interface {
	Create(ctx context.Context, krsAction *KrsAction) error
	FindAll(ctx context.Context) ([]*KrsAction, error)
	FindByID(ctx context.Context, id string) (*KrsAction, error)
	Update(ctx context.Context, krsAction *KrsAction) error
	Delete(ctx context.Context, id string) error
}
