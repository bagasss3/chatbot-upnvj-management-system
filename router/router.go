package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group          *echo.Group
	userController model.UserController
	authController model.AuthController
}

func NewRouter(group *echo.Group, userController model.UserController, authController model.AuthController) {
	rt := &router{
		group:          group,
		userController: userController,
		authController: authController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
	r.authRouter()
}
