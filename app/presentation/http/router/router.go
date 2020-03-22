package router

import (
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) {
	e.POST("/projects", c.CreateProject)
	e.GET("/projects", c.GetProjects)
	e.GET("/projects/:id", c.GetProject)
	e.PUT("/projects/:id", c.UpdateProject)
	e.DELETE("/projects/:id", c.DeleteProject)
}
