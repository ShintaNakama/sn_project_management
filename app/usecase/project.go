package usecase

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

// Project UseCase interface

type ProjectUseCase interface {
	GetProjects(ctx context.Context) ([]*model.Project, error)
	GetProject(ctx context.Context, id int) (*model.Project, error)
	CreateProject(ctx context.Context, project *model.Project) (*model.Project, error)
	UpdateProject(ctx context.Context, project *model.Project, id int) (*model.Project, error)
	DeleteProject(ctx context.Context, id int) error
}

type projectUseCase struct {
	repository.ProjectRepository
}

// NewProjectUseCase return projectUseCase
func NewProjectUseCase(r repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{r}
}

func (u *projectUseCase) GetProjects(ctx context.Context) ([]*model.Project, error) {
	return u.ProjectRepository.Fetch(ctx)
}

func (u *projectUseCase) GetProject(ctx context.Context, id int) (*model.Project, error) {
	return u.ProjectRepository.FetchByID(ctx, id)
}

func (u *projectUseCase) CreateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	id, err := u.ProjectRepository.Create(ctx, project)
	if err != nil {
		return nil, err
	}
	return u.ProjectRepository.FetchByID(ctx, id)
}

func (u *projectUseCase) UpdateProject(ctx context.Context, project *model.Project, id int) (*model.Project, error) {
	if err := u.ProjectRepository.Update(ctx, project, id); err != nil {
		return project, err
	}
	return u.ProjectRepository.FetchByID(ctx, id)
}

func (u *projectUseCase) DeleteProject(ctx context.Context, id int) error {
	return u.ProjectRepository.Delete(ctx, id)
}
