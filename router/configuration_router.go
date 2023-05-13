package router

import "cbupnvj/middleware"

func (r *router) configurationRouter() {
	r.group.POST("/configuration", r.configurationController.HandleCreateConfiguration(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/configuration", r.configurationController.HandleFindAllConfiguration(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/configuration/:id", r.configurationController.HandleFindConfiguration(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/configuration/:id", r.configurationController.HandleUpdateConfiguration(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())

	r.group.POST("/training-history", r.trainingHistoryController.HandleCreateTrainingHistory(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/training-history", r.trainingHistoryController.HandleFindAllTrainingHistory(), middleware.MustAuthenticateAccessToken())
}
