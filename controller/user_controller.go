package controller

import (
	"cbupnvj/constant"
	"cbupnvj/model"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type userController struct {
	userService model.UserService
}

func NewUserController(userService model.UserService) model.UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) HandleCreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.CreateAdminRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		create, err := u.userService.CreateAdmin(c.Request().Context(), model.CreateAdminRequest{
			Email:      req.Email,
			Name:       req.Name,
			MajorId:    req.MajorId,
			Password:   req.Password,
			Repassword: req.Repassword,
			Type:       req.Type,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, create)
	}
}
