package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group               *echo.Group
	userController      model.UserController
	authController      model.AuthController
	intentController    model.IntentController
	utteranceController model.UtteranceController
}

func NewRouter(group *echo.Group, userController model.UserController,
	authController model.AuthController,
	intentController model.IntentController,
	utteranceController model.UtteranceController) {
	rt := &router{
		group:               group,
		userController:      userController,
		authController:      authController,
		intentController:    intentController,
		utteranceController: utteranceController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
	r.authRouter()
	r.intentRouter()
	r.utteranceRouter()
}
