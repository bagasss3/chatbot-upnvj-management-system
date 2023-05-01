package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReqBodyDataType string
type HttpMethod string

const (
	DataTypeString ReqBodyDataType = "STRING"
	DataTypeInt    ReqBodyDataType = "INT"
	DataTypeFloat  ReqBodyDataType = "FLOAT"
	DataTypeDate   ReqBodyDataType = "DATE"
)

const (
	HttpMethodPost HttpMethod = "POST"
	HttpMethodPut  HttpMethod = "PUT"
)

type CreateReqBodyActionRequest struct {
	ReqName  string          `json:"req_name" validate:"required,min=3,max=60"`
	DataType ReqBodyDataType `json:"data_type" validate:"required"`
}

func (c *CreateReqBodyActionRequest) Validate() error {
	return validate.Struct(c)
}

type CreateReqBodyActionArrayRequest struct {
	ActionHttpId string                        `json:"action_http_id" validate:"required"`
	PostFields   []*CreateReqBodyActionRequest `json:"post_fields" validate:"omitempty"`
	PutFields    []*CreateReqBodyActionRequest `json:"put_fields" validate:"omitempty"`
}

func (c *CreateReqBodyActionArrayRequest) Validate() error {
	return validate.Struct(c)
}

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
	Id           string          `json:"id"`
	ActionHttpId string          `json:"action_http_id"`
	ReqName      string          `json:"req_name"`
	DataType     ReqBodyDataType `json:"data_type"`
	Method       HttpMethod      `json:"method"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type ReqBodyController interface {
	HandleCreateReqBody() echo.HandlerFunc
	HandleFindAllReqBodyByActionHttpID() echo.HandlerFunc
	HandleUpdateReqBody() echo.HandlerFunc
	HandleDeleteReqBody() echo.HandlerFunc
}

type ReqBodyService interface {
	CreateReqBody(ctx context.Context, req CreateReqBodyActionArrayRequest) (bool, error)
	FindAllReqBodyByActionHttpID(ctx context.Context, actionHttpID string, method string) ([]*ReqBody, error)
	FindByID(ctx context.Context, id string) (*ReqBody, error)
	UpdateReqBody(ctx context.Context, id, actionHttpID string, req UpdateReqBodyRequest) (*ReqBody, error)
	DeleteReqBody(ctx context.Context, id, actionHttpID string) (bool, error)
}

type ReqBodyRepository interface {
	Create(ctx context.Context, tx *gorm.DB, reqBody *ReqBody) error
	FindAll(ctx context.Context, actionHttpID string, method HttpMethod) ([]*ReqBody, error)
	FindByID(ctx context.Context, id string) (*ReqBody, error)
	Update(ctx context.Context, id string, reqBody *ReqBody) error
	Delete(ctx context.Context, id string) error
}
