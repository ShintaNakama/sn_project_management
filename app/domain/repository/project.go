package repository

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
)

type ProjectRepository interface {
	Fetch(ctx context.Context) ([]*model.Project, error)
	FetchByID(ctx context.Context, id int) (*model.Project, error)
	Create(ctx context.Context, project *model.Project) (int, error)
	Update(ctx context.Context, project *model.Project, id int) error
	Delete(ctx context.Context, id int) error
}
