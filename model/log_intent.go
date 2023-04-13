package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateUpdateLogIntentRequest struct {
	IntentId int64 `json:"intent_id" validate:"required"`
}

func (c *CreateUpdateLogIntentRequest) Validate() error {
	return validate.Struct(c)
}

type LogIntent struct {
	Id        int64     `json:"id"`
	IntentId  int64     `json:"intent_id"`
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
	FindLogIntentByIntentID(ctx context.Context, intentId int64) (*LogIntent, error)
	FindAllLogIntent(ctx context.Context) ([]*LogIntent, error)
}

type LogIntentRepository interface {
	Create(ctx context.Context, li *LogIntent) error
	FindByIntentID(ctx context.Context, intentId int64) (*LogIntent, error)
	FindAll(ctx context.Context) ([]*LogIntent, error)
	Update(ctx context.Context, intentId int64, li *LogIntent) error
}
