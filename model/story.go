package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUpdateStoryRequest struct {
	StoryTitle string `json:"story_title" validate:"required,min=3,max=60"`
}

func (c *CreateUpdateStoryRequest) Validate() error {
	return validate.Struct(c)
}

type Story struct {
	Id         string         `json:"id"`
	StoryTitle string         `json:"story_title"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type StoryController interface {
	HandleCreateStory() echo.HandlerFunc
	HandleFindAllStory() echo.HandlerFunc
	HandleFindStoryByID() echo.HandlerFunc
	HandleUpdateStory() echo.HandlerFunc
	HandleDeleteStory() echo.HandlerFunc
}

type StoryService interface {
	CreateStory(ctx context.Context, req CreateUpdateStoryRequest) (*Story, error)
	FindAllStory(ctx context.Context) ([]*Story, error)
	FindStoryByID(ctx context.Context, id string) (*Story, error)
	UpdateStory(ctx context.Context, id string, req CreateUpdateStoryRequest) (*Story, error)
	DeleteStory(ctx context.Context, id string) (bool, error)
}

type StoryRepository interface {
	Create(ctx context.Context, story *Story) error
	FindAll(ctx context.Context) ([]*Story, error)
	FindByID(ctx context.Context, id string) (*Story, error)
	Update(ctx context.Context, id string, Story *Story) error
	Delete(ctx context.Context, id string) error
}
