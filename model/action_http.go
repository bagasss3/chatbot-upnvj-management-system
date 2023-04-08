package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateActionHttpRequest struct {
	ActionId     int64  `json:"action_id" validate:"required"`
	GetHttpReq   string `json:"get_http_req" validate:"required"`
	PostHttpReq  string `json:"post_http_req" validate:"omitempty"`
	PutHttpReq   string `json:"put_http_req" validate:"omitempty"`
	DelHttpReq   string `json:"del_http_req" validate:"omitempty"`
	ApiKey       string `json:"api_key" validate:"omitempty"`
	TextResponse string `json:"text_response" validate:"required,min=3,max=100"`
}

func (c *CreateActionHttpRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateActionHttpRequest struct {
	GetHttpReq   string `json:"get_http_req" validate:"required"`
	PostHttpReq  string `json:"post_http_req" validate:"omitempty"`
	PutHttpReq   string `json:"put_http_req" validate:"omitempty"`
	DelHttpReq   string `json:"del_http_req" validate:"omitempty"`
	ApiKey       string `json:"api_key" validate:"omitempty"`
	TextResponse string `json:"text_response" validate:"required,min=3,max=100"`
}

func (c *UpdateActionHttpRequest) Validate() error {
	return validate.Struct(c)
}

type ActionHttp struct {
	Id           int64          `json:"id"`
	ActionId     int64          `json:"action_id"`
	GetHttpReq   string         `json:"get_http_req"`
	PostHttpReq  string         `json:"post_http_req"`
	PutHttpReq   string         `json:"put_http_req"`
	DelHttpReq   string         `json:"del_http_req"`
	ApiKey       string         `json:"api_key"`
	TextResponse string         `json:"text_response"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type ActionHttpController interface {
	HandleCreateActionHttp() echo.HandlerFunc
	HandleFindActionHttpByActionID() echo.HandlerFunc
	HandleUpdateActionHttp() echo.HandlerFunc
	HandleDeleteActionHttp() echo.HandlerFunc
}

type ActionHttpService interface {
	CreateActionHttp(ctx context.Context, req CreateActionHttpRequest) (*ActionHttp, error)
	FindActionHttpByID(ctx context.Context, actionId int64) (*ActionHttp, error)
	UpdateActionHttp(ctx context.Context, actionId int64, req UpdateActionHttpRequest) (*ActionHttp, error)
	DeleteActionHttp(ctx context.Context, actionId int64) (bool, error)
}

type ActionHttpRepository interface {
	Create(ctx context.Context, actionHttp *ActionHttp) error
	FindByID(ctx context.Context, id int64) (*ActionHttp, error)
	FindByActionID(ctx context.Context, actionId int64) (*ActionHttp, error)
	Update(ctx context.Context, actionHttp *ActionHttp) error
	Delete(ctx context.Context, actionId int64) error
}
