package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	Id                    int64          `json:"id"`
	AccessToken           string         `json:"access_token"`
	RefreshToken          string         `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time      `json:"refresh_token_expired_at"`
	UserID                int64          `json:"user_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}
