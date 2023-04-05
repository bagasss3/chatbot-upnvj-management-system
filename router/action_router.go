package router

import (
	"cbupnvj/middleware"
)

func (r *router) actionRouter() {
	r.group.POST("/action", r.actionController.HandleCreateAction(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action", r.actionController.HandleFindAllAction(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/:id", r.actionController.HandleFindActionByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/:id", r.actionController.HandleUpdateAction(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/:id", r.actionController.HandleDeleteAction(), middleware.MustAuthenticateAccessToken())
}
