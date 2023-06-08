package router

import (
	"cbupnvj/middleware"
)

func (r *router) intentRouter() {
	r.group.POST("/intent", r.intentController.HandleCreateIntent(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent", r.intentController.HandleFindAllIntent())
	r.group.GET("/intent/example", r.intentController.HandleFindAllWithExamples())
	r.group.GET("/intent/:id", r.intentController.HandleFindIntentByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/intent/:id", r.intentController.HandleUpdateIntent(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/intent/:id", r.intentController.HandleDeleteIntent(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/intent/entity", r.entityController.HandleCreateEntity(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/:intentId/entity", r.entityController.HandleFindAllEntity(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/:intentId/entity/:id", r.entityController.HandleFindEntityByID(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/intent/:intentId/entity/:id", r.entityController.HandleDeleteEntity(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/dashboard/log/intent", r.logIntentController.HandleCreateOrUpdateLogIntent(), middleware.MustHaveAPIKey())
	r.group.GET("/dashboard/log/intent", r.logIntentController.HandleFindAllLogIntent(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/dashboard/log/intent/:intentId", r.logIntentController.HandleFindLogIntentByIntentID(), middleware.MustAuthenticateAccessToken())

	r.group.GET("/intent/count", r.intentController.HandleCountAllIntent(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/intent/fallback", r.fallbackChatLogController.HandleCreateFallbackChatLog(), middleware.MustHaveAPIKey())
	r.group.GET("/intent/fallback", r.fallbackChatLogController.HandleFindAllFallbackChatLog(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/log/fallback", r.fallbackChatLogController.HandleFindAllFallbackChatLogGroupCluster(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/intent/all/fallback", r.fallbackChatLogController.HandleFindAllFallbackChatLogOldAndNew(), middleware.MustHaveAPIKey())
	r.group.GET("/intent/fallback/update", r.fallbackChatLogController.HandleUpdateGroupCluster(), middleware.MustAuthenticateAccessToken())
}
