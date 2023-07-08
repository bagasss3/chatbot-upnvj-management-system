package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateIntentRequest struct {
	Name                  string `json:"name" validate:"required,min=3,max=60"`
	IsInformationAcademic bool   `json:"is_information_academic" validate:"required"`
}

func (c *CreateUpdateIntentRequest) Validate() error {
	return validate.Struct(c)
}

type Intent struct {
	Id                    string         `json:"id"`
	Name                  string         `json:"name"`
	IsInformationAcademic bool           `json:"is_information_academic"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
	Examples              []*Example     `json:"examples" gorm:"foreignKey:IntentId"`
}

type IntentController interface {
	HandleCreateIntent() echo.HandlerFunc
	HandleFindAllIntent() echo.HandlerFunc
	HandleFindIntentByID() echo.HandlerFunc
	HandleUpdateIntent() echo.HandlerFunc
	HandleDeleteIntent() echo.HandlerFunc
	HandleCountAllIntent() echo.HandlerFunc
	HandleFindAllWithExamples() echo.HandlerFunc
	HandleFindAllInformationAcademics() echo.HandlerFunc
}

type IntentService interface {
	CreateIntent(ctx context.Context, req CreateUpdateIntentRequest) (*Intent, error)
	FindAllIntent(ctx context.Context, name string) ([]*Intent, error)
	FindIntentByID(ctx context.Context, id string) (*Intent, error)
	UpdateIntent(ctx context.Context, id string, req CreateUpdateIntentRequest) (*Intent, error)
	DeleteIntent(ctx context.Context, id string) (bool, error)
	CountAllIntent(ctx context.Context) (int64, error)
	FindAllWithExamples(ctx context.Context) ([]*Intent, error)
	FindAllInformationAcademics(ctx context.Context) ([]*Intent, error)
}

type IntentRepository interface {
	Create(ctx context.Context, intent *Intent) error
	FindByID(ctx context.Context, id string) (*Intent, error)
	FindByName(ctx context.Context, id string) (*Intent, error)
	FindAll(ctx context.Context, name string) ([]*Intent, error)
	Update(ctx context.Context, id string, intent *Intent) error
	DeleteWithTx(ctx context.Context, id string, tx *gorm.DB) error
	CountAll(ctx context.Context) (int64, error)
	FindAllWithExamples(ctx context.Context) ([]*Intent, error)
	FindAllInformationAcademics(ctx context.Context) ([]*Intent, error)
}
