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

	r.group.POST("/action/http", r.actionHttpController.HandleCreateActionHttp(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/:actionId/http", r.actionHttpController.HandleFindActionHttpByActionID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/:actionId/http", r.actionHttpController.HandleUpdateActionHttp(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/:actionId/http", r.actionHttpController.HandleDeleteActionHttp(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/action/http/req", r.reqBodyController.HandleCreateReqBody(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/http/:actionHttpId/req", r.reqBodyController.HandleFindAllReqBodyByActionHttpID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/http/:actionHttpId/req/:reqBodyId", r.reqBodyController.HandleUpdateReqBody(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/http/:actionHttpId/req/:reqBodyId", r.reqBodyController.HandleDeleteReqBody(), middleware.MustAuthenticateAccessToken())
}
