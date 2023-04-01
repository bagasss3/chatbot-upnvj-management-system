package constant

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// http errors
var (
	ErrInvalidArgument     = echo.NewHTTPError(http.StatusBadRequest, "invalid argument")
	ErrNotFound            = echo.NewHTTPError(http.StatusNotFound, "record not found")
	ErrInternal            = echo.NewHTTPError(http.StatusInternalServerError, "internal system error")
	ErrUnauthenticated     = echo.NewHTTPError(http.StatusUnauthorized, "unauthenticated")
	ErrUnauthorized        = echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	ErrRefreshTokenExpired = echo.NewHTTPError(http.StatusNotAcceptable, "refresh token expired")
	ErrPasswordMismatch    = echo.NewHTTPError(http.StatusBadRequest, "password mismatch")
)
