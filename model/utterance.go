package model

import (
	"time"

	"gorm.io/gorm"
)

type Utterance struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	Response  string         `json:"response"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
