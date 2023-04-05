package router

import (
	"cbupnvj/middleware"
)

func (r *router) exampleRouter() {
	r.group.POST("/example", r.exampleController.HandleCreateExample(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/example/:id", r.exampleController.HandleFindAllExample(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/example/:intentId/:exampleId", r.exampleController.HandleFindExampleByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/example/:intentId/:exampleId", r.exampleController.HandleUpdateExample(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/example/:intentId/:exampleId", r.exampleController.HandleDeleteExample(), middleware.MustAuthenticateAccessToken())
}
