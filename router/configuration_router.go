package router

import "cbupnvj/middleware"

func (r *router) configurationRouter() {
	r.group.POST("/configuration", r.configurationController.HandleCreateConfiguration(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/configuration/:id", r.configurationController.HandleFindConfiguration(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/configuration/:id", r.configurationController.HandleUpdateConfiguration(), middleware.MustAuthenticateAccessToken())
}
