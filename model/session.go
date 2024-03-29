package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	Id                    string         `json:"id"`
	AccessToken           string         `json:"access_token"`
	AccessTokenExpiredAt  time.Time      `json:"access_token_expired_at"`
	RefreshToken          string         `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time      `json:"refresh_token_expired_at"`
	UserID                string         `json:"user_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at"`
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) error
	FindByID(ctx context.Context, id string) (*Session, error)
	FindByRefreshToken(ctx context.Context, refreshToken string) (*Session, error)
	RefreshToken(ctx context.Context, session *Session) error
	Delete(ctx context.Context, userId string) error
}
