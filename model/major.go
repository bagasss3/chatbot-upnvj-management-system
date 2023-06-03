package model

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Major struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	FacultyId string `json:"faculty_id"`
}

type MajorController interface {
	HandleFindAllMajor() echo.HandlerFunc
	HandleFindByIDMajor() echo.HandlerFunc
}

type MajorService interface {
	FindAllMajor(ctx context.Context) ([]*Major, error)
	FindByIDMajor(ctx context.Context, id string) (*Major, error)
}

type MajorRepository interface {
	FindAll(ctx context.Context) ([]*Major, error)
	FindByID(ctx context.Context, id string) (*Major, error)
}
