package router

import (
	"cbupnvj/middleware"
)

func (r *router) conversationRouter() {
	r.group.POST("/rule", r.ruleController.HandleCreateRule(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/rule", r.ruleController.HandleFindAllRule(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/rule/:id", r.ruleController.HandleFindRuleByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/rule/:id", r.ruleController.HandleUpdateRule(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/rule/:id", r.ruleController.HandleDeleteRule(), middleware.MustAuthenticateAccessToken())
}
