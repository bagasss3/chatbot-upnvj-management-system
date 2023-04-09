package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type StepType string

const (
	StepIntent    StepType = "INTENT"
	StepUtterance StepType = "UTTERANCE"
	StepAction    StepType = "ACTION_HTTP"
)

type CreateStepRequest struct {
	StoryId    int64    `json:"story_id" validate:"required"`
	ResponseId int64    `json:"response_id" validate:"required"`
	Type       StepType `json:"type" validate:"required"`
}

func (c *CreateStepRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateStepRequest struct {
	ResponseId int64    `json:"response_id" validate:"required"`
	Type       StepType `json:"type" validate:"required"`
}

func (c *UpdateStepRequest) Validate() error {
	return validate.Struct(c)
}

type Step struct {
	Id         int64     `json:"id"`
	StoryId    int64     `json:"story_id"`
	Type       StepType  `json:"type"`
	ResponseId int64     `json:"response_id"`
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
	FindAllStepByStoryID(ctx context.Context, storyId int64) ([]*Step, error)
	FindStepByID(ctx context.Context, id, storyId int64) (*Step, error)
	UpdateStep(ctx context.Context, id, storyId int64, req UpdateStepRequest) (*Step, error)
	DeleteStep(ctx context.Context, id, storyId int64) (bool, error)
}

type StepRepository interface {
	Create(ctx context.Context, step *Step) error
	FindAll(ctx context.Context, storyId int64) ([]*Step, error)
	FindByID(ctx context.Context, id, storyId int64) (*Step, error)
	Update(ctx context.Context, id int64, step *Step) error
	Delete(ctx context.Context, id int64) error
	DeleteAllByStoryID(ctx context.Context, storyId int64) error
}
