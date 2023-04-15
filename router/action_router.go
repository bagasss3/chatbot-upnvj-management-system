package router

import (
	"cbupnvj/middleware"
)

func (r *router) actionRouter() {
	r.group.POST("/action/http", r.actionHttpController.HandleCreateActionHttp(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/http", r.actionHttpController.HandleFindAllActionHttp(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/http/:id", r.actionHttpController.HandleFindActionHttpByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/http/:id", r.actionHttpController.HandleUpdateActionHttp(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/http/:id", r.actionHttpController.HandleDeleteActionHttp(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/action/http/req", r.reqBodyController.HandleCreateReqBody(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/http/:actionHttpId/req", r.reqBodyController.HandleFindAllReqBodyByActionHttpID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/http/:actionHttpId/req/:reqBodyId", r.reqBodyController.HandleUpdateReqBody(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/http/:actionHttpId/req/:reqBodyId", r.reqBodyController.HandleDeleteReqBody(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/action/krs", r.krsActionController.HandleCreateKrsAction(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/krs", r.krsActionController.HandleFindAllKrsAction(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/action/krs/:id", r.krsActionController.HandleFindKrsActionByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/action/krs/:id", r.krsActionController.HandleUpdateKrsAction(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/action/krs/:id", r.krsActionController.HandleDeleteKrsAction(), middleware.MustAuthenticateAccessToken())

	r.group.GET("/action/http/count", r.actionHttpController.HandleCountAllActionHttp(), middleware.MustAuthenticateAccessToken())
}
