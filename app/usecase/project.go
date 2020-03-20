package usecase

import "context"

// Project UseCase interface

type ProjectUseCase interface {
	List(ctx context.Context) ([]*model.Project, error)
	Show(ctx context.Context, id int) (*model.Project, error)
	Create(ctx context.Context, project *model.Project) (*model.Project, error)
	Update(ctx context.Context, id int) (*model.Project, error)
	Delete(ctx context.Context, id int) error
}

type projectUseCase struct {
	repository.ProjectRepository
}

// NewprojectUseCase projectUseCaseを取得します.
func NewProjectUseCase(r repository.ProjectRepository) projectUseCase {
	return &projectUseCase{r}
}

func (u *projectUseCase) GetProjects(ctx context.Context) ([]*model.Project, error) {
	return u.ProjectRepository.Fetch(ctx)
}

func (u *projectUseCase) GetProject(ctx context.Context, id int) (*model.Project, error) {
	return u.ProjectRepository.FetchByID(ctx, id)
}

func (u *projectUseCase) CreateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	return u.ProjectRepository.Create(ctx, project)
}

func (u *projectUseCase) UpdateProject(ctx context.Context, id int) (*model.Project, error) {
	project, err := u.ProjectRepository.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.ProjectRepository.Update(ctx, project)
}

func (u *projectUseCase) DeleteProject(ctx context.Context, id int) error {
	return u.ProjectRepository.Delete(ctx, id)
}
