package model

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserType string

const (
	UserAdmin      UserType = "ADMIN"
	UserSuperAdmin UserType = "SUPER_ADMIN"
)

type CreateAdminRequest struct {
	Email      string
	Name       string
	Type       UserType
	MajorId    int64
	Password   string
	Repassword string
}

type User struct {
	Id        int64          `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	Type      UserType       `json:"type"`
	MajorId   int64          `json:"major_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserController interface {
	HandleCreateAdmin() echo.HandlerFunc
}

type UserService interface {
	CreateAdmin(ctx context.Context, req CreateAdminRequest) (*User, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, userEmail string) (*User, error)
	ResetPassword(ctx context.Context, user *User) error
}
