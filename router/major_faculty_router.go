package router

func (r *router) majorFacultyRouter() {
	r.group.GET("/majors", r.majorController.HandleFindAllMajor())
	r.group.GET("/majors/:id", r.majorController.HandleFindByIDMajor())

	r.group.GET("/faculties", r.facultyController.HandleFindAllFaculty())
	r.group.GET("/faculties/:id", r.facultyController.HandleFindByIDFaculty())
}
