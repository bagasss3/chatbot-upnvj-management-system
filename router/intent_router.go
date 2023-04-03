package router

import (
	"cbupnvj/middleware"
)

func (r *router) intentRouter() {
	r.group.POST("/intent", r.intentController.HandleCreateIntent(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent", r.intentController.HandleFindAllIntent(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/:id", r.intentController.HandleFindIntentByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/intent/:id", r.intentController.HandleUpdateIntent(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/intent/:id", r.intentController.HandleDeleteIntent(), middleware.MustAuthenticateAccessToken())
}
