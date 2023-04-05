package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateActionRequest struct {
	Name string `json:"name" validate:"required,min=3,max=30"`
}

func (c *CreateUpdateActionRequest) Validate() error {
	return validate.Struct(c)
}

type Action struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ActionController interface {
	HandleCreateAction() echo.HandlerFunc
	HandleFindAllAction() echo.HandlerFunc
	HandleFindActionByID() echo.HandlerFunc
	HandleUpdateAction() echo.HandlerFunc
	HandleDeleteAction() echo.HandlerFunc
}

type ActionService interface {
	CreateAction(ctx context.Context, req CreateUpdateActionRequest) (*Action, error)
	FindAllAction(ctx context.Context) ([]*Action, error)
	FindActionByID(ctx context.Context, id int64) (*Action, error)
	UpdateAction(ctx context.Context, id int64, req CreateUpdateActionRequest) (*Action, error)
	DeleteAction(ctx context.Context, id int64) (bool, error)
}

type ActionRepository interface {
	Create(ctx context.Context, action *Action) error
	FindAll(ctx context.Context) ([]*Action, error)
	FindByID(ctx context.Context, id int64) (*Action, error)
	Update(ctx context.Context, id int64, action *Action) error
	Delete(ctx context.Context, id int64) error
}
