package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group                *echo.Group
	userController       model.UserController
	authController       model.AuthController
	intentController     model.IntentController
	utteranceController  model.UtteranceController
	exampleController    model.ExampleController
	actionController     model.ActionController
	actionHttpController model.ActionHttpController
}

func NewRouter(group *echo.Group, userController model.UserController,
	authController model.AuthController,
	intentController model.IntentController,
	utteranceController model.UtteranceController,
	exampleController model.ExampleController,
	actionController model.ActionController,
	actionHttpController model.ActionHttpController) {
	rt := &router{
		group:                group,
		userController:       userController,
		authController:       authController,
		intentController:     intentController,
		utteranceController:  utteranceController,
		exampleController:    exampleController,
		actionController:     actionController,
		actionHttpController: actionHttpController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
	r.authRouter()
	r.intentRouter()
	r.utteranceRouter()
	r.exampleRouter()
	r.actionRouter()
}
