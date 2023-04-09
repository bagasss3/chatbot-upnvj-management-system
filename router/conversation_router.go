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

	r.group.POST("/story", r.storyController.HandleCreateStory(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/story", r.storyController.HandleFindAllStory(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/story/:id", r.storyController.HandleFindStoryByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/story/:id", r.storyController.HandleUpdateStory(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/story/:id", r.storyController.HandleDeleteStory(), middleware.MustAuthenticateAccessToken())

	r.group.POST("/story/step", r.stepController.HandleCreateStep(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/story/:storyId/step", r.stepController.HandleFindAllStepByStoryID(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/story/:storyId/step/:id", r.stepController.HandleFindStepByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/story/:storyId/step/:id", r.stepController.HandleUpdateStep(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/story/:storyId/step/:id", r.stepController.HandleDeleteStep(), middleware.MustAuthenticateAccessToken())
}
