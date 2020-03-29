package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/usecase"
	"github.com/labstack/echo"
)

// ProjectController is CRUD
type ProjectController interface {
	GetProjects(c echo.Context) error
	GetProject(c echo.Context) error
	CreateProject(c echo.Context) error
	UpdateProject(c echo.Context) error
	DeleteProject(c echo.Context) error
}

type projectController struct {
	ProjectUseCase usecase.ProjectUseCase
}

// NewProjectController is return projectController(project handler)
func NewProjectController(u usecase.ProjectUseCase) ProjectController {
	return &projectController{u}
}

// Get Projects is Project List
func (p *projectController) GetProjects(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	projects, err := p.ProjectUseCase.GetProjects(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Projects does not exist.")
	}
	return c.JSON(http.StatusOK, projects)
}

// Get Project is Project
func (p *projectController) GetProject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	project, err := p.ProjectUseCase.GetProject(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Project does not exist.")
	}
	return c.JSON(http.StatusOK, project)
}

// Create Project
func (p *projectController) CreateProject(c echo.Context) error {
	project := &model.Project{}
	if err := c.Bind(project); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	//project, err := p.ProjectUseCase.CreateProject(ctx, project)
	err := p.ProjectUseCase.CreateProject(ctx, project)
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError, "Project can not Create.")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, project)
}

// Update Project
func (p *projectController) UpdateProject(c echo.Context) error {
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	// return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
	//}
	project := &model.Project{}
	if err := c.Bind(project); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	project, err := p.ProjectUseCase.UpdateProject(ctx, project)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Project can not Update.")
	}
	return c.JSON(http.StatusOK, project)
}

// Delete Project
func (p *projectController) DeleteProject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	if err := p.ProjectUseCase.DeleteProject(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Project can not Delete.")
	}
	return c.NoContent(http.StatusNoContent)
}
