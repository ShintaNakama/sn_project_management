// project/controller test
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
type mockProjectUseCase struct{}

func (u *mockProjectUseCase) GetProjects(ctx context.Context) ([]*model.Project, error) {
	return getMockProjects(5), nil
}

func (u *mockProjectUseCase) GetProject(ctx context.Context, id int) (*model.Project, error) {
	return getMockProject(id), nil
}
func (u *mockProjectUseCase) CreateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	return getMockProject(1), nil
}

func (u *mockProjectUseCase) UpdateProject(ctx context.Context, project *model.Project, id int) (*model.Project, error) {
	mp := getMockProject(id)
	mp.Name = mp.Name + "_updated"
	return mp, nil
}

func (u *mockProjectUseCase) DeleteProject(ctx context.Context, id int) error {
	return nil
}

func getMockProjects(n int) []*model.Project {
	projects := []*model.Project{}
	for i := 1; i < n; i++ {
		p := getMockProject(i)
		projects = append(projects, p)
	}
	return projects
}

func getMockProject(n int) *model.Project {
	p := &model.Project{
		ID:          n,
		Name:        fmt.Sprintf("test_project_name_%d", n),
		Description: fmt.Sprintf("test_project_descriptin_%d", n),
		CreatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		UpdatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		Completed:   0,
	}
	return p
}

func getMockProjectNoID() *model.Project {
	u := &model.Project{
		Name:        fmt.Sprintf("test_project_name_%d", 1),
		Description: fmt.Sprintf("test_project_descriptin_%d", 1),
		CreatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
		UpdatedAt:   time.Date(2020, 4, 1, 12, 35, 42, 123456789, time.Local),
	}
	return u
}

// GetProjects
func TestControllerGetProjects(t *testing.T) {
	// expected
	expected := getMockProjects(5)

	// set stub
	usecase := &mockProjectUseCase{}
	controller := NewProjectController(usecase)

	e := echo.New()
	req := httptest.NewRequest("GET", "/projects", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controller.GetProjects(c)) {
		projects := []*model.Project{}
		if err := json.Unmarshal(rec.Body.Bytes(), &projects); err != nil {
			t.Fatal(err)
		}
		t.Log(rec.Code)
		assert.Equal(t, http.StatusOK, rec.Code)

		for i, p := range projects {
			t.Log(p)
			assert.Equal(t, expected[i], p)
		}
	}
}

// Table Driven Tests
type getProjectTest struct {
	ID      int
	Project *model.Project
}

var getProjectTests = []getProjectTest{
	{math.MaxInt8, getMockProject(math.MaxInt8)},
	{math.MaxInt16, getMockProject(math.MaxInt16)},
	{math.MaxInt32, getMockProject(math.MaxInt32)},
	{math.MaxInt64, getMockProject(math.MaxInt64)},
}

// GetProject
func TestControllerGetProject(t *testing.T) {
	// set stub
	usecase := &mockProjectUseCase{}
	controller := NewProjectController(usecase)

	for _, test := range getProjectTests {
		e := echo.New()
		req := httptest.NewRequest("GET", "/projects", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/projects/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		if assert.NoError(t, controller.GetProject(c)) {
			project := &model.Project{}
			if err := json.Unmarshal(rec.Body.Bytes(), &project); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			t.Log(project)
			assert.Equal(t, test.Project, project)
		}
	}
}

// Create Project
func TestControllerCreateProject(t *testing.T) {
	// expected
	expected := getMockProject(1)

	// set stub
	usecase := &mockProjectUseCase{}
	h := NewProjectController(usecase)

	e := echo.New()
	jsonBytes, err := json.Marshal(getMockProjectNoID())
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(echo.POST, "/projects", strings.NewReader(string(jsonBytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, h.CreateProject(c)) {
		project := &model.Project{}
		if err := json.Unmarshal(rec.Body.Bytes(), &project); err != nil {
			t.Fatal(err)
		}

		t.Log(rec.Code)
		assert.Equal(t, http.StatusCreated, rec.Code)
		t.Log(project)
		assert.Equal(t, expected, project)
	}
}

// Update Project
type updateProjectTest struct {
	ID          int
	ProjectName string
}

var updateProjectTests = []updateProjectTest{
	{math.MaxInt8, fmt.Sprintf("test_project_name_%d_updated", math.MaxInt8)},
	{math.MaxInt16, fmt.Sprintf("test_project_name_%d_updated", math.MaxInt16)},
	{math.MaxInt32, fmt.Sprintf("test_project_name_%d_updated", math.MaxInt32)},
	{math.MaxInt64, fmt.Sprintf("test_project_name_%d_updated", math.MaxInt64)},
}

func TestControllerUpdateproject(t *testing.T) {
	// set stub
	usecase := &mockProjectUseCase{}
	ctl := NewProjectController(usecase)

	for _, test := range updateProjectTests {
		// set request
		e := echo.New()
		jsonBytes, err := json.Marshal(test)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(echo.PUT, "/projects", strings.NewReader(string(jsonBytes)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/projects/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		// assertions
		if assert.NoError(t, ctl.UpdateProject(c)) {
			project := &model.Project{}
			if err := json.Unmarshal(rec.Body.Bytes(), &project); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			t.Log(project)
			assert.Equal(t, test.ProjectName, project.Name)
		}
	}
}

// Delete Project
func TestControllerDeleteProject(t *testing.T) {
	// set stub
	usecase := &mockProjectUseCase{}
	h := NewProjectController(usecase)

	// set request
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/projects", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/projects/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(1))

	// assertions
	if assert.NoError(t, h.DeleteProject(c)) {
		t.Log(rec.Code)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
