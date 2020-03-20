package repository

type ProjectRepository interface {
  Fetch(ctx context.Context) ([]*model.Project, error)
	FetchByID(ctx context.Context, id int) (*model.Project, error)
	Create(ctx context.Context, user *model.Project) (*model.Project, error)
	Update(ctx context.Context, user *model.Project) (*model.Project, error)
	Delete(ctx context.Context, id int) error
}
