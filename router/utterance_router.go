package router

import (
	"cbupnvj/middleware"
)

func (r *router) utteranceRouter() {
	r.group.POST("/utterance", r.utteranceController.HandleCreateUtterance(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/utterance", r.utteranceController.HandleFindAllUtterance(), middleware.MustAuthenticateAccessToken())
	r.group.GET("/utterance/:id", r.utteranceController.HandleFindUtteranceByID(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/utterance/:id", r.utteranceController.HandleUpdateUtterance(), middleware.MustAuthenticateAccessToken())
	r.group.DELETE("/utterance/:id", r.utteranceController.HandleDeleteUtterance(), middleware.MustAuthenticateAccessToken())
}
