package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

// DB database interface
type projectRepository struct {
	Conn *sql.DB
}

// NewProjectRepository is return projectRepository
func NewProjectRepository(Conn *sql.DB) repository.ProjectRepository {
	return &projectRepository{Conn}
}

func (r *projectRepository) Fetch(ctx context.Context) ([]*model.Project, error) {
	projects := make([]*model.Project, 0)
	rows, err := r.Conn.QueryContext(ctx, "select id, name, description, created_at, updated_at, completed from projects")
	if err != nil {
		return projects, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		p := &model.Project{}
		if err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt, &p.Completed); err != nil {
			return projects, err
		}
		projects = append(projects, p)
	}
	fmt.Println(projects)

	return projects, nil
}

func (r *projectRepository) FetchByID(ctx context.Context, id int) (*model.Project, error) {
	project := &model.Project{}
	if err := r.Conn.QueryRowContext(
		ctx,
		"select * from projects where id=?",
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.Completed,
	); err != nil {
		return project, err
	}

	return project, nil
}

func (r *projectRepository) Create(ctx context.Context, p *model.Project) (*model.Project, error) {
	//tx, err := conn.Begin()
	//if err != nil {
	//	return p, err
	//}
	result, err := r.Conn.ExecContext(
		ctx,
		"insert into projects (name, description, created_at, updated_at) values (?,?,NOW(),NOW())",
		p.Name,
		p.Description,
	)
	if err != nil {
		return p, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return p, err
	}
	project := &model.Project{}
	if err = r.Conn.QueryRowContext(
		ctx,
		"select * from projects where id=?",
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.Completed,
	); err != nil {
		return project, err
	}
	//if err = tx.Commit(); err != nil {
	//	return p, err
	//}
	return project, nil
}

// Update
func (r *projectRepository) Update(ctx context.Context, p *model.Project, id int) (*model.Project, error) {
	result, err := r.Conn.ExecContext(
		ctx,
		"update projects set name=?, description=?, updated_at=NOW() where id=?",
		p.Name,
		p.Description,
		id,
	)
	if err != nil {
		return p, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return p, err
	}
	if rows != 1 {
		err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
		return p, err
	}
	project := &model.Project{}
	if err = r.Conn.QueryRowContext(
		ctx,
		"select * from projects where id=?",
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.Completed,
	); err != nil {
		return project, err
	}
	//if err = tx.Commit(); err != nil {
	//	return p, err
	//}
	return project, nil
}

// Delete
func (r *projectRepository) Delete(ctx context.Context, id int) error {
	result, err := r.Conn.ExecContext(
		ctx,
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
