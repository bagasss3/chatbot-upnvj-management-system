package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateUtteranceRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=60"`
	Response string `json:"response" validate:"required,min=3,max=200"`
}

func (c *CreateUpdateUtteranceRequest) Validate() error {
	return validate.Struct(c)
}

type Utterance struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Response  string         `json:"response"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UtteranceController interface {
	HandleCreateUtterance() echo.HandlerFunc
	HandleFindAllUtterance() echo.HandlerFunc
	HandleFindUtteranceByID() echo.HandlerFunc
	HandleUpdateUtterance() echo.HandlerFunc
	HandleDeleteUtterance() echo.HandlerFunc
	HandleCountAllUtterance() echo.HandlerFunc
}

type UtteranceService interface {
	CreateUtterance(ctx context.Context, req CreateUpdateUtteranceRequest) (*Utterance, error)
	FindAllUtterance(ctx context.Context) ([]*Utterance, error)
	FindUtteranceByID(ctx context.Context, id string) (*Utterance, error)
	UpdateUtterance(ctx context.Context, id string, req CreateUpdateUtteranceRequest) (*Utterance, error)
	DeleteUtterance(ctx context.Context, id string) (bool, error)
	CountAllUtterance(ctx context.Context) (int64, error)
}

type UtteranceRepository interface {
	Create(ctx context.Context, utterance *Utterance) error
	FindByID(ctx context.Context, id string) (*Utterance, error)
	FindAll(ctx context.Context) ([]*Utterance, error)
	Update(ctx context.Context, id string, utterance *Utterance) error
	Delete(ctx context.Context, id string) error
	CountAll(ctx context.Context) (int64, error)
}
