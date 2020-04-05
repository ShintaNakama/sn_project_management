package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/usecase"
	"github.com/labstack/echo"
)

// TaskController is CRUD
type TaskController interface {
	GetTasks(c echo.Context) error
	GetTask(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	TaskUseCase usecase.TaskUseCase
}

// NewTaskController is return taskController(task handler)
func NewTaskController(u usecase.TaskUseCase) TaskController {
	return &taskController{u}
}

// Get Tasks is Task List
func (p *taskController) GetTasks(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("project_id"))
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	tasks, err := p.TaskUseCase.GetTasks(ctx, pID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Tasks does not exist.")
	}
	return c.JSON(http.StatusOK, tasks)
}

// Get Task is Task
func (p *taskController) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Task ID must be int")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	task, err := p.TaskUseCase.GetTask(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Task does not exist.")
	}
	return c.JSON(http.StatusOK, task)
}

// Create Task
func (p *taskController) CreateTask(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("project_id"))
  //pID, err := strconv.Atoi(c.FormValue("project_id"))
  if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Project ID must be int")
  }

	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}
  log.Println(task)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	task, err = p.TaskUseCase.CreateTask(ctx, task, pID)
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError, "Task can not Create.")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, task)
}

// Update Task
func (p *taskController) UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Task ID must be int")
  }
	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	task, err = p.TaskUseCase.UpdateTask(ctx, task, id)
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError, "Task can not Update.")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, task)
}

// Delete Task
func (p *taskController) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Task ID must be int")
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	if err := p.TaskUseCase.DeleteTask(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Task can not Delete.")
	}
	return c.NoContent(http.StatusNoContent)
}
