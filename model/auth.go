package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email, PlainPassword string
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type AuthController interface {
	HandleLoginByEmailAndPassword() echo.HandlerFunc
	HandleRefreshToken() echo.HandlerFunc
}

type AuthService interface {
	LoginByEmailAndPassword(ctx context.Context, req LoginRequest) (*Session, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest) (*Session, error)
}
