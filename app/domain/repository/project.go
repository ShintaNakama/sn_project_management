package repository

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
)

type ProjectRepository interface {
	Fetch(ctx context.Context) ([]*model.Project, error)
	FetchByID(ctx context.Context, id int) (*model.Project, error)
	Create(ctx context.Context, project *model.Project) (*model.Project, error)
	Update(ctx context.Context, project *model.Project, id int) (*model.Project, error)
	Delete(ctx context.Context, id int) error
}
