package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateExampleRequest struct {
	IntentID int64  `json:"intent_id" validate:"required"`
	Example  string `json:"example" validate:"required,min=3,max=200"`
}

func (c *CreateExampleRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateExampleRequest struct {
	Example string `json:"example" validate:"required,min=3,max=200"`
}

func (c *UpdateExampleRequest) Validate() error {
	return validate.Struct(c)
}

type Example struct {
	Id        int64     `json:"id"`
	IntentId  int64     `json:"intent_id"`
	Example   string    `json:"example"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ExampleController interface {
	HandleCreateExample() echo.HandlerFunc
	HandleFindAllExample() echo.HandlerFunc
	HandleFindExampleByID() echo.HandlerFunc
	HandleUpdateExample() echo.HandlerFunc
	HandleDeleteExample() echo.HandlerFunc
}

type ExampleService interface {
	CreateExample(ctx context.Context, req CreateExampleRequest) (*Example, error)
	FindAllExampleByIntentID(ctx context.Context, id int64) ([]*Example, error)
	FindExampleByIntentID(ctx context.Context, intentId, exampleId int64) (*Example, error)
	UpdateExample(ctx context.Context, intentId, exampleId int64, req UpdateExampleRequest) (*Example, error)
	DeleteExample(ctx context.Context, intentId, exampleId int64) (bool, error)
}

type ExampleRepository interface {
	Create(ctx context.Context, example *Example) error
	FindAllByIntentID(ctx context.Context, intentId int64) ([]*Example, error)
	FindByID(ctx context.Context, intentId, exampleId int64) (*Example, error)
	Update(ctx context.Context, id int64, example *Example) error
	Delete(ctx context.Context, id int64) error
}
