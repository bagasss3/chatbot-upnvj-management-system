package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateActionHttpRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=60"`
	GetHttpReq   string `json:"get_http_req" validate:"required"`
	PostHttpReq  string `json:"post_http_req" validate:"omitempty"`
	PutHttpReq   string `json:"put_http_req" validate:"omitempty"`
	DelHttpReq   string `json:"del_http_req" validate:"omitempty"`
	ApiKey       string `json:"api_key" validate:"omitempty"`
	TextResponse string `json:"text_response" validate:"required,min=3,max=100"`
}

func (c *CreateUpdateActionHttpRequest) Validate() error {
	return validate.Struct(c)
}

type ActionHttp struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
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
	HandleFindAllActionHttp() echo.HandlerFunc
	HandleFindActionHttpByID() echo.HandlerFunc
	HandleUpdateActionHttp() echo.HandlerFunc
	HandleDeleteActionHttp() echo.HandlerFunc
	HandleCountAllActionHttp() echo.HandlerFunc
}

type ActionHttpService interface {
	CreateActionHttp(ctx context.Context, req CreateUpdateActionHttpRequest) (*ActionHttp, error)
	FindAllActionHttp(ctx context.Context, name string) ([]*ActionHttp, error)
	FindActionHttpByID(ctx context.Context, id string) (*ActionHttp, error)
	UpdateActionHttp(ctx context.Context, id string, req CreateUpdateActionHttpRequest) (*ActionHttp, error)
	DeleteActionHttp(ctx context.Context, id string) (bool, error)
	CountAllActionHttp(ctx context.Context) (int64, error)
}

type ActionHttpRepository interface {
	Create(ctx context.Context, actionHttp *ActionHttp) error
	FindAll(ctx context.Context, name string) ([]*ActionHttp, error)
	FindByID(ctx context.Context, id string) (*ActionHttp, error)
	Update(ctx context.Context, actionHttp *ActionHttp) error
	Delete(ctx context.Context, id string) error
	CountAll(ctx context.Context) (int64, error)
}
