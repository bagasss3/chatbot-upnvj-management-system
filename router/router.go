package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group            *echo.Group
	userController   model.UserController
	authController   model.AuthController
	intentController model.IntentController
}

func NewRouter(group *echo.Group, userController model.UserController,
	authController model.AuthController,
	intentController model.IntentController) {
	rt := &router{
		group:            group,
		userController:   userController,
		authController:   authController,
		intentController: intentController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
	r.authRouter()
	r.intentRouter()
}
