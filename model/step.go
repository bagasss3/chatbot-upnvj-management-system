package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StepType string

const (
	StepIntent    StepType = "INTENT"
	StepUtterance StepType = "UTTERANCE"
	StepAction    StepType = "ACTION_HTTP"
)

type CreateStepRequest struct {
	StoryId    string   `json:"story_id" validate:"required"`
	ResponseId string   `json:"response_id" validate:"required"`
	Type       StepType `json:"type" validate:"required"`
}

func (c *CreateStepRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateStepRequest struct {
	ResponseId string   `json:"response_id" validate:"required"`
	Type       StepType `json:"type" validate:"required"`
}

func (c *UpdateStepRequest) Validate() error {
	return validate.Struct(c)
}

type Step struct {
	Id         string    `json:"id"`
	StoryId    string    `json:"story_id"`
	Type       StepType  `json:"type"`
	ResponseId string    `json:"response_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StepController interface {
	HandleCreateStep() echo.HandlerFunc
	HandleFindAllStepByStoryID() echo.HandlerFunc
	HandleFindStepByID() echo.HandlerFunc
	HandleUpdateStep() echo.HandlerFunc
	HandleDeleteStep() echo.HandlerFunc
}

type StepService interface {
	CreateStep(ctx context.Context, req CreateStepRequest) (*Step, error)
	FindAllStepByStoryID(ctx context.Context, storyId string) ([]*Step, error)
	FindStepByID(ctx context.Context, id, storyId string) (*Step, error)
	UpdateStep(ctx context.Context, id, storyId string, req UpdateStepRequest) (*Step, error)
	DeleteStep(ctx context.Context, id, storyId string) (bool, error)
}

type StepRepository interface {
	Create(ctx context.Context, step *Step) error
	FindAll(ctx context.Context, storyId string) ([]*Step, error)
	FindByID(ctx context.Context, id, storyId string) (*Step, error)
	Update(ctx context.Context, id string, step *Step) error
	Delete(ctx context.Context, id string) error
	DeleteAllByStoryID(ctx context.Context, tx *gorm.DB, storyId string) error
}
