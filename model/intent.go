package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateIntentRequest struct {
	Name string `json:"name" validate:"required,min=3,max=60"`
}

func (c *CreateUpdateIntentRequest) Validate() error {
	return validate.Struct(c)
}

type Intent struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type IntentController interface {
	HandleCreateIntent() echo.HandlerFunc
	HandleFindAllIntent() echo.HandlerFunc
	HandleFindIntentByID() echo.HandlerFunc
	HandleUpdateIntent() echo.HandlerFunc
	HandleDeleteIntent() echo.HandlerFunc
	HandleCountAllIntent() echo.HandlerFunc
}

type IntentService interface {
	CreateIntent(ctx context.Context, req CreateUpdateIntentRequest) (*Intent, error)
	FindAllIntent(ctx context.Context) ([]*Intent, error)
	FindIntentByID(ctx context.Context, id int64) (*Intent, error)
	UpdateIntent(ctx context.Context, id int64, req CreateUpdateIntentRequest) (*Intent, error)
	DeleteIntent(ctx context.Context, id int64) (bool, error)
	CountAllIntent(ctx context.Context) (int64, error)
}

type IntentRepository interface {
	Create(ctx context.Context, intent *Intent) error
	FindByID(ctx context.Context, id int64) (*Intent, error)
	FindAll(ctx context.Context) ([]*Intent, error)
	Update(ctx context.Context, id int64, intent *Intent) error
	Delete(ctx context.Context, id int64) error
	CountAll(ctx context.Context) (int64, error)
}
