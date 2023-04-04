package model

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type TokenType string

const (
	ACCESS_TOKEN_TYPE  TokenType = "ACCESS_TOKEN"
	REFRESH_TOKEN_TYPE TokenType = "REFRESH_TOKEN"
)

type Claims struct {
	UserID int64     `json:"userID"`
	Role   UserType  `json:"role"`
	Type   TokenType `json:"type"`
	jwt.StandardClaims
}

type UserAuth struct {
	UserID int64    `json:"userID"`
	Role   UserType `json:"role"`
}

type LoginRequest struct {
	Email, Password string
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type ForgotPasswordRequest struct {
	Email string
}

type AuthController interface {
	HandleLoginByEmailAndPassword() echo.HandlerFunc
	HandleRefreshToken() echo.HandlerFunc
	HandleForgotPassword() echo.HandlerFunc
}

type AuthService interface {
	LoginByEmailAndPassword(ctx context.Context, req LoginRequest) (*Session, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest) (*Session, error)
	ForgotPassword(ctx context.Context, req ForgotPasswordRequest) (bool, error)
}
