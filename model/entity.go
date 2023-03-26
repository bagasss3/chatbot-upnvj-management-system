package model

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	IntentId  int64          `json:"intent_id"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
