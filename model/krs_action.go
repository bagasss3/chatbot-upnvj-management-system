package model

import (
	"time"

	"gorm.io/gorm"
)

type KrsAction struct {
	Id         int64          `json:"id"`
	ActionId   int64          `json:"action_id"`
	GetHttpReq string         `json:"get_http_req"`
	ApiKey     string         `json:"api_key"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
