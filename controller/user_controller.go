package controller

import (
	"cbupnvj/constant"
	"cbupnvj/middleware"
	"cbupnvj/model"
	"net/http"
	"strconv"

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

func (u *userController) HandleFindAllAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := u.userService.FindAllAdmin(c.Request().Context())
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, users)
	}
}

func (u *userController) HandleFindAdminByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		user, err := u.userService.FindAdminByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (u *userController) HandleUpdateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.UpdateAdminRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		update, err := u.userService.UpdateAdmin(c.Request().Context(), id, model.UpdateAdminRequest{
			Name:       req.Name,
			MajorId:    req.MajorId,
			Password:   req.Password,
			Repassword: req.Repassword,
		})
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}

func (u *userController) HandleDeleteAdminByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")

		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Error(err)
			return err
		}

		isDeleted, err := u.userService.DeleteAdminByID(c.Request().Context(), id)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, isDeleted)
	}
}

func (u *userController) HandleProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		ctxUser := middleware.GetUserFromCtx(ctx)

		user, err := u.userService.FindAdminByID(ctx, ctxUser.UserID)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (u *userController) HandleUpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := model.UpdateUserPasswordRequest{}
		if err := c.Bind(&req); err != nil {
			log.Error(err)
			return constant.ErrInternal
		}

		ctx := c.Request().Context()
		ctxUser := middleware.GetUserFromCtx(ctx)

		update, err := u.userService.UpdateProfile(c.Request().Context(), ctxUser.UserID, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, update)
	}
}
