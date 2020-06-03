package router

import (
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) {
	// project
	e.POST("/projects", c.CreateProject)
	e.GET("/projects", c.GetProjects)
	e.GET("/projects/:id", c.GetProject)
	e.PUT("/projects/:id", c.UpdateProject)
	e.DELETE("/projects/:id", c.DeleteProject)
	// task
	e.POST("/projects/:project_id/tasks", c.CreateTask)
	e.GET("/projects/:project_id/tasks", c.GetTasks)
	e.GET("/projects/:project_id/tasks/:id", c.GetTask)
	e.PUT("/projects/:project_id/tasks/:id", c.UpdateTask)
	e.DELETE("/projects/:project_id/tasks/:id", c.DeleteTask)
}
