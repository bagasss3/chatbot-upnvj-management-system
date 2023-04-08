package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateEntityRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	IntentId int64  `json:"intent_id" validate:"required"`
}

func (c *CreateEntityRequest) Validate() error {
	return validate.Struct(c)
}

type Entity struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	IntentId  int64          `json:"intent_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type EntityController interface {
	HandleCreateEntity() echo.HandlerFunc
	HandleFindAllEntity() echo.HandlerFunc
	HandleFindEntityByID() echo.HandlerFunc
	HandleDeleteEntity() echo.HandlerFunc
}

type EntityService interface {
	CreateEntity(ctx context.Context, req CreateEntityRequest) (*Entity, error)
	FindAllEntity(ctx context.Context, intentId int64) ([]*Entity, error)
	FindEntityByID(ctx context.Context, id, intentId int64) (*Entity, error)
	DeleteEntity(ctx context.Context, id, intentId int64) (bool, error)
}

type EntityRepository interface {
	Create(ctx context.Context, entity *Entity) error
	FindAll(ctx context.Context, intentId int64) ([]*Entity, error)
	FindByID(ctx context.Context, id int64) (*Entity, error)
	Delete(ctx context.Context, id int64) error
}
