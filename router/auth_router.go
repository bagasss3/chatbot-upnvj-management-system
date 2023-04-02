package router

func (r *router) authRouter() {
	r.group.POST("/login", r.authController.HandleLoginByEmailAndPassword())
	r.group.POST("/forgot-password", r.authController.HandleForgotPassword())
	r.group.POST("/refresh-token", r.authController.HandleRefreshToken())
}
