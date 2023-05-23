package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateUpdateLogIntentRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateUpdateLogIntentRequest) Validate() error {
	return validate.Struct(c)
}

type LogIntent struct {
	Id        string    `json:"id"`
	IntentId  string    `json:"intent_id"`
	Intent    Intent    `json:"intent"`
	Mention   int       `json:"mention"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LogIntentController interface {
	HandleCreateOrUpdateLogIntent() echo.HandlerFunc
	HandleFindLogIntentByIntentID() echo.HandlerFunc
	HandleFindAllLogIntent() echo.HandlerFunc
}

type LogIntentService interface {
	CreateOrUpdateLogIntent(ctx context.Context, req CreateUpdateLogIntentRequest) (*LogIntent, error)
	FindLogIntentByIntentID(ctx context.Context, intentId string) (*LogIntent, error)
	FindAllLogIntent(ctx context.Context) ([]*LogIntent, error)
}

type LogIntentRepository interface {
	Create(ctx context.Context, li *LogIntent) error
	FindByIntentID(ctx context.Context, intentId string) (*LogIntent, error)
	FindAll(ctx context.Context) ([]*LogIntent, error)
	Update(ctx context.Context, intentId string, li *LogIntent) error
}
