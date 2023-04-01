package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group          *echo.Group
	userController model.UserController
}

func NewRouter(group *echo.Group, userController model.UserController) {
	rt := &router{
		group:          group,
		userController: userController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
}
