package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type authController struct {
	authService model.AuthService
}

func NewAuthController(authService model.AuthService) model.AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) HandleLoginByEmailAndPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.LoginRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		login, err := a.authService.LoginByEmailAndPassword(c.Request().Context(), model.LoginRequest{
			Email:         req.Email,
			PlainPassword: req.PlainPassword,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, login)
	}
}

func (a *authController) HandleRefreshToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.RefreshTokenRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		refreshToken, err := a.authService.RefreshToken(c.Request().Context(), model.RefreshTokenRequest{
			RefreshToken: req.RefreshToken,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, refreshToken)
	}
}
