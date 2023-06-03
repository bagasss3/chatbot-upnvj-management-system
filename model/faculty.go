package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Faculty struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Majors []*Major `json:"majors" gorm:"foreignKey:FacultyId"`
}

type FacultyController interface {
	HandleFindAllFaculty() echo.HandlerFunc
	HandleFindByIDFaculty() echo.HandlerFunc
}

type FacultyService interface {
	FindAllFaculty(ctx context.Context) ([]*Faculty, error)
	FindByIDFaculty(ctx context.Context, id string) (*Faculty, error)
}

type FacultyRepository interface {
	FindAll(ctx context.Context) ([]*Faculty, error)
	FindByID(ctx context.Context, id string) (*Faculty, error)
}
