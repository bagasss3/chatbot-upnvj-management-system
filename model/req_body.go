package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateReqBodyRequest struct {
	ActionHttpId string `json:"action_http_id" validate:"required"`
	ReqName      string `json:"req_name" validate:"required,min=3,max=60"`
}

func (c *CreateReqBodyRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateReqBodyRequest struct {
	ReqName string `json:"req_name" validate:"required,min=3,max=60"`
}

func (c *UpdateReqBodyRequest) Validate() error {
	return validate.Struct(c)
}

type ReqBody struct {
	Id           string    `json:"id"`
	ActionHttpId string    `json:"action_http_id"`
	ReqName      string    `json:"req_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ReqBodyController interface {
	HandleCreateReqBody() echo.HandlerFunc
	HandleFindAllReqBodyByActionHttpID() echo.HandlerFunc
	HandleUpdateReqBody() echo.HandlerFunc
	HandleDeleteReqBody() echo.HandlerFunc
}

type ReqBodyService interface {
	CreateReqBody(ctx context.Context, req CreateReqBodyRequest) (*ReqBody, error)
	FindAllReqBodyByActionHttpID(ctx context.Context, actionHttpID string) ([]*ReqBody, error)
	FindByID(ctx context.Context, id string) (*ReqBody, error)
	UpdateReqBody(ctx context.Context, id, actionHttpID string, req UpdateReqBodyRequest) (*ReqBody, error)
	DeleteReqBody(ctx context.Context, id, actionHttpID string) (bool, error)
}

type ReqBodyRepository interface {
	Create(ctx context.Context, reqBody *ReqBody) error
	FindAll(ctx context.Context, actionHttpID string) ([]*ReqBody, error)
	FindByID(ctx context.Context, id string) (*ReqBody, error)
	Update(ctx context.Context, id string, reqBody *ReqBody) error
	Delete(ctx context.Context, id string) error
}
