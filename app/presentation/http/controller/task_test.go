// task/controller test
package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// cotrollerのテストなので、controllrの各methodが使うusercaseをmockで定義する
type mockTaskUseCase struct{}

func (u *mockTaskUseCase) GetTasks(ctx context.Context, pID int) ([]*model.Task, error) {
	return getMockTasks(pID, 3), nil
}

func (u *mockTaskUseCase) GetTask(ctx context.Context, id int) (*model.Task, error) {
	return getMockTask(id, 1), nil
}
func (u *mockTaskUseCase) CreateTask(ctx context.Context, task *model.Task, pID int) (*model.Task, error) {
	return getMockTask(1, pID), nil
}

func (u *mockTaskUseCase) UpdateTask(ctx context.Context, task *model.Task, id int) (*model.Task, error) {
	mp := getMockTask(id, 1)
	mp.Description = mp.Description + "_updated"
	return mp, nil
}

func (u *mockTaskUseCase) DeleteTask(ctx context.Context, id int) error {
	return nil
}

func getMockTasks(pID, n int) []*model.Task {
	tasks := []*model.Task{}
	for tID := 1; tID < n; tID++ {
		t := getMockTask(tID, pID)
		tasks = append(tasks, t)
	}
	return tasks
}

func getMockTask(tID, pID int) *model.Task {
	p := &model.Task{
		ID:          tID,
		ProjectID:   pID,
		Description: fmt.Sprintf("test_task_description_%d", tID),
		CreatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		UpdatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		Completed:   0,
	}
	return p
}

func getMockTaskNoID(pID int) *model.Task {
	u := &model.Task{
		ProjectID:   pID,
		Description: fmt.Sprintf("test_task_description_%d", 1),
		CreatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		UpdatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
	}
	return u
}

// Table Driven Tests
type getTasksTest struct {
	ProjectID int
	Tasks     []*model.Task
}

var getTasksTests = []getTasksTest{
	{1, getMockTasks(1, 3)},
	{2, getMockTasks(2, 3)},
}

// GetTasks
func TestControllerGetTasks(t *testing.T) {

	// set stub
	usecase := &mockTaskUseCase{}
	controller := NewTaskController(usecase)

	for _, test := range getTasksTests {
		e := echo.New()
		req := httptest.NewRequest("GET", "/projects/tasks", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/projects/:project_id")
		c.SetParamNames("project_id")
		c.SetParamValues(fmt.Sprint(test.ProjectID))

		if assert.NoError(t, controller.GetTasks(c)) {
			tasks := []*model.Task{}
			if err := json.Unmarshal(rec.Body.Bytes(), &tasks); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)

			for i, task := range tasks {
				t.Log(task)
				assert.Equal(t, test.Tasks[i], task)
			}
		}
	}
}

// Table Driven Tests
type getTaskTest struct {
	ID        int
	ProjectID int
	Task      *model.Task
}

var getTaskTests = []getTaskTest{
	{math.MaxInt8, 1, getMockTask(math.MaxInt8, 1)},
	{math.MaxInt16, 1, getMockTask(math.MaxInt16, 1)},
}

// GetTask
func TestControllerGetTask(t *testing.T) {
	// set stub
	usecase := &mockTaskUseCase{}
	controller := NewTaskController(usecase)

	for _, test := range getTaskTests {
		e := echo.New()
		req := httptest.NewRequest("GET", "/projects/tasks", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/projects/:project_id")
		c.SetParamNames("project_id")
		c.SetParamValues(fmt.Sprint(test.ProjectID))
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		if assert.NoError(t, controller.GetTask(c)) {
			task := &model.Task{}
			if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, test.Task, task)
		}
	}
}

// Create Task
func TestControllerCreateTask(t *testing.T) {
	// expected
	pID := 1
	expected := getMockTask(pID, 1)

	// set stub
	usecase := &mockTaskUseCase{}
	h := NewTaskController(usecase)

	e := echo.New()
	jsonBytes, err := json.Marshal(getMockTaskNoID(pID))
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(echo.POST, "/projects/tasks", strings.NewReader(string(jsonBytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/projects/:project_id")
	c.SetParamNames("project_id")
	c.SetParamValues(fmt.Sprint(pID))
	// Assertions
	if assert.NoError(t, h.CreateTask(c)) {
		task := &model.Task{}
		if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
			t.Fatal(err)
		}

		t.Log(rec.Code)
		assert.Equal(t, http.StatusCreated, rec.Code)
		t.Log(task)
		assert.Equal(t, expected, task)
	}
}

// Update Task
type updateTaskTest struct {
	ID          int
	ProjectID          int
	TaskDescription string
}

var updateTaskTests = []updateTaskTest{
	{math.MaxInt8, 1,fmt.Sprintf("test_task_description_%d_updated", math.MaxInt8)},
	{math.MaxInt16, 1,fmt.Sprintf("test_task_description_%d_updated", math.MaxInt16)},
}

func TestControllerUpdateTask(t *testing.T) {
	// set stub
	usecase := &mockTaskUseCase{}
	ctl := NewTaskController(usecase)

	for _, test := range updateTaskTests {
		// set request
		e := echo.New()
		jsonBytes, err := json.Marshal(test)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(echo.PUT, "/projects/tasks", strings.NewReader(string(jsonBytes)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/projects/:project_id")
		c.SetParamNames("project_id")
		c.SetParamValues(fmt.Sprint(test.ProjectID))
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		// assertions
		if assert.NoError(t, ctl.UpdateTask(c)) {
			task := &model.Task{}
			if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			t.Log(task)
			assert.Equal(t, test.TaskDescription, task.Description)
		}
	}
}

// Delete Task
func TestControllerDeleteTask(t *testing.T) {
	// set stub
	usecase := &mockTaskUseCase{}
	h := NewTaskController(usecase)

	// set request
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/projects/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
  c.SetPath("/projects/:project_id")
  c.SetParamNames("project_id")
  c.SetParamValues(fmt.Sprint(1))
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(1))

	// assertions
	if assert.NoError(t, h.DeleteTask(c)) {
		t.Log(rec.Code)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
