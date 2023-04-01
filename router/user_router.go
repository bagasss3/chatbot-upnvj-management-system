package router

func (r *router) userRouter() {
	r.group.POST("/user", r.userController.HandleCreateAdmin())
}
