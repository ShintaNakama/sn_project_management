package repository

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
)

type TaskRepository interface {
	Fetch(ctx context.Context, pID int) ([]*model.Task, error)
	FetchByID(ctx context.Context, id int) (*model.Task, error)
	Create(ctx context.Context, task *model.Task, pID int) (int, error)
	Update(ctx context.Context, task *model.Task, id int) error
	Delete(ctx context.Context, id int) error
}
