package router

import (
	"cbupnvj/middleware"
)

func (r *router) userRouter() {
	r.group.POST("/user", r.userController.HandleCreateAdmin(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())
	r.group.GET("/user", r.userController.HandleFindAllAdmin(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())
	r.group.GET("/user/:id", r.userController.HandleFindAdminByID(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())
	r.group.PUT("/user/:id", r.userController.HandleUpdateAdmin(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())
	r.group.DELETE("/user/:id", r.userController.HandleDeleteAdminByID(), middleware.MustAuthenticateAccessToken(), middleware.MustSuperAdminOnly())
	r.group.GET("/user/profile", r.userController.HandleProfile(), middleware.MustAuthenticateAccessToken())
	r.group.PUT("/user/profile", r.userController.HandleUpdateProfile(), middleware.MustAuthenticateAccessToken())
}
