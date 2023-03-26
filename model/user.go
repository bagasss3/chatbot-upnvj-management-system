package model

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	UserAdmin      UserType = "ADMIN"
	UserSuperAdmin UserType = "SUPER_ADMIN"
)

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
