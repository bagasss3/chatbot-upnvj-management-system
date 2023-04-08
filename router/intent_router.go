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

	r.group.POST("/intent/entity", r.entityController.HandleCreateEntity(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/:intentId/entity", r.entityController.HandleFindAllEntity(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/:intentId/entity/:id", r.entityController.HandleFindEntityByID(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/intent/:intentId/entity/:id", r.entityController.HandleDeleteEntity(), middleware.MustAuthenticateAccessToken())
}
