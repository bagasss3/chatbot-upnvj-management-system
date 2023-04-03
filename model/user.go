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
	Email      string   `json:"email" validate:"required,email"`
	Name       string   `json:"name" validate:"required,min=3,max=60"`
	Type       UserType `json:"type" validate:"required"`
	MajorId    int64    `json:"major_id" validate:"required"`
	Password   string   `json:"password" validate:"required"`
	Repassword string   `json:"repassword" validate:"required"`
}

func (c *CreateAdminRequest) Validate() error {
	return validate.Struct(c)
}

type UpdateAdminRequest struct {
	Name       string `json:"name" validate:"required,min=3,max=60"`
	MajorId    int64  `json:"major_id" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Repassword string `json:"repassword" validate:"required"`
}

func (c *UpdateAdminRequest) Validate() error {
	return validate.Struct(c)
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
	HandleFindAllAdmin() echo.HandlerFunc
	HandleFindAdminByID() echo.HandlerFunc
	HandleUpdateAdmin() echo.HandlerFunc
	HandleDeleteAdminByID() echo.HandlerFunc
	HandleProfile() echo.HandlerFunc
	HandleUpdateProfile() echo.HandlerFunc
}

type UserService interface {
	CreateAdmin(ctx context.Context, req CreateAdminRequest) (*User, error)
	FindAllAdmin(ctx context.Context) ([]*User, error)
	FindAdminByID(ctx context.Context, id int64) (*User, error)
	UpdateAdmin(ctx context.Context, id int64, req UpdateAdminRequest) (*User, error)
	DeleteAdminByID(ctx context.Context, id int64) (bool, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, userEmail string) (*User, error)
	ResetPassword(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, id int64, user *User) error
	Delete(ctx context.Context, id int64) error
}
