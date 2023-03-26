package model

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	Id         int64          `json:"id"`
	StoryTitle string         `json:"story_title"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
