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
type projectRepository struct {
	Conn *gorp.DbMap
}

// NewProjectRepository is return projectRepository
func NewProjectRepository(Conn *gorp.DbMap) repository.ProjectRepository {
	return &projectRepository{Conn}
}

func (r *projectRepository) Fetch(ctx context.Context) ([]*model.Project, error) {
	projects := make([]*model.Project, 0)
	_, err := r.Conn.Select(&projects, "select id, name, description, created_at, updated_at, completed from projects")
	if err != nil {
		return projects, err
	}
	return projects, nil
}

func (r *projectRepository) FetchByID(ctx context.Context, id int) (*model.Project, error) {
	project := &model.Project{}
	err := r.Conn.SelectOne(&project, "select * from projects where id=?", id)
	if err != nil {
		return project, err
	}

	return project, nil
}

// Create
func (r *projectRepository) Create(ctx context.Context, p *model.Project) (int, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	err := r.Conn.Insert(p)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var id int
	err = r.Conn.SelectOne(&id, "select id from projects order by id desc limit 1")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

// Update
func (r *projectRepository) Update(ctx context.Context, p *model.Project, id int) error {
	// Start a new transaction
	trans, err := r.Conn.Begin()
	if err != nil {
		return err
	}
	project := &model.Project{}
	err = trans.SelectOne(&project, "select * from projects where id=?", id)
	if err != nil {
		return err
	}
	project.Name = p.Name
	project.Description = p.Description
	project.UpdatedAt = time.Now()
	rows, err := trans.Update(project)
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
func (r *projectRepository) Delete(ctx context.Context, id int) error {
	result, err := r.Conn.Exec(
		"delete from projects where id=?",
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
