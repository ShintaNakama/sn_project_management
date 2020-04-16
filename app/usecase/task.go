package usecase

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

// Task UseCase interface

type TaskUseCase interface {
	GetTasks(ctx context.Context, pID int) ([]*model.Task, error)
	GetTask(ctx context.Context, id int) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task, pID int) (*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task, id int) (*model.Task, error)
	DeleteTask(ctx context.Context, id int) error
}

type taskUseCase struct {
	repository.TaskRepository
}

// NewTaskUseCase return taskUseCase
func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{r}
}

func (u *taskUseCase) GetTasks(ctx context.Context, pID int) ([]*model.Task, error) {
	return u.TaskRepository.Fetch(ctx, pID)
}

func (u *taskUseCase) GetTask(ctx context.Context, id int) (*model.Task, error) {
	return u.TaskRepository.FetchByID(ctx, id)
}

func (u *taskUseCase) CreateTask(ctx context.Context, task *model.Task, pID int) (*model.Task, error) {
	id, err := u.TaskRepository.Create(ctx, task, pID)
	if err != nil {
		return nil, err
	}
	return u.TaskRepository.FetchByID(ctx, id)
}

func (u *taskUseCase) UpdateTask(ctx context.Context, task *model.Task, id int) (*model.Task, error) {
	if err := u.TaskRepository.Update(ctx, task, id); err != nil {
		return task, err
	}
	return u.TaskRepository.FetchByID(ctx, id)
}

func (u *taskUseCase) DeleteTask(ctx context.Context, id int) error {
	return u.TaskRepository.Delete(ctx, id)
}
