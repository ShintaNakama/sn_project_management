package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
	"github.com/go-gorp/gorp"
)

// DB database interface
type taskRepository struct {
	Conn *gorp.DbMap
}

// NewTaskRepository is return taskRepository
func NewTaskRepository(Conn *gorp.DbMap) repository.TaskRepository {
	return &taskRepository{Conn}
}

func (r *taskRepository) Fetch(ctx context.Context, pID int) ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)
	_, err := r.Conn.Select(&tasks, "select id, project_id, description, created_at, updated_at, completed from tasks where project_id=?", pID)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepository) FetchByID(ctx context.Context, id int) (*model.Task, error) {
	task := &model.Task{}
	err := r.Conn.SelectOne(&task, "select * from tasks where id=?", id)
	if err != nil {
		return task, err
	}

	return task, nil
}

// Create
func (r *taskRepository) Create(ctx context.Context, t *model.Task, pID int) (int, error) {
  t.ProjectID = pID
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
  err := r.Conn.Insert(t)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var id int
	err = r.Conn.SelectOne(&id, "select id from tasks order by id desc limit 1")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

// Update
func (r *taskRepository) Update(ctx context.Context, t *model.Task, id int) error {
  // Start a new transaction
  trans, err := r.Conn.Begin()
  if err != nil {
      return err
  }
	task := &model.Task{}
	err = trans.SelectOne(&task, "select * from tasks where id=?", id)
	if err != nil {
		return err
	}
	task.Description = t.Description
	task.UpdatedAt = time.Now()
	rows, err := trans.Update(task)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows != 1 {
		err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
		return err
	}
  // if the commit is successful, a nil error is returned
  if err = trans.Commit(); err != nil {
		log.Println(err)
		return err
  }
	return nil
}

// Delete
func (r *taskRepository) Delete(ctx context.Context, id int) error {
	result, err := r.Conn.Exec(
		"delete from tasks where id=?",
		id,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
		return err
	}
	return err
}
