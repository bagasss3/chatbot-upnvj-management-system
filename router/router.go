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
	actionHttpController model.ActionHttpController
	reqBodyController    model.ReqBodyController
	krsActionController  model.KrsActionController
}

func NewRouter(group *echo.Group, userController model.UserController,
	authController model.AuthController,
	intentController model.IntentController,
	utteranceController model.UtteranceController,
	exampleController model.ExampleController,
	actionHttpController model.ActionHttpController,
	reqBodyController model.ReqBodyController,
	krsActionController model.KrsActionController) {
	rt := &router{
		group:                group,
		userController:       userController,
		authController:       authController,
		intentController:     intentController,
		utteranceController:  utteranceController,
		exampleController:    exampleController,
		actionHttpController: actionHttpController,
		reqBodyController:    reqBodyController,
		krsActionController:  krsActionController,
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
