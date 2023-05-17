package router

import (
	"cbupnvj/middleware"
)

func (r *router) workerRouter() {
	r.group.GET("/train", r.workerController.HandleStartTrainingModel(), middleware.MustAuthenticateAccessToken())
}
