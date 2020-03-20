package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/ShintaNakama/sn_project_management/app/domain/model"
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

//const dsn = "root@tcp(db)/sn_project_management"

// DB database interface
type projectRepository struct {
	Conn *sql.DB
}

func NewProjectRepository(Conn *sql.DB) repository.ProjectRepository {
	return &projectRepository{Conn}
}

func (r *projectRepository) Fetch(ctx context.Context) ([]*model.Project, error) {
	projects := make([]*model.Project, 0)
	rows, err := r.Conn.QueryContext(
		ctx,
		"SELECT id, name, description, created_at, updated_at, completed FROM projects",
	)
	if err != nil {
		return projects, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		var p *model.Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt, p.UpdatedAt, p.Completed); err != nil {
			return projects, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *projectRepository) FetchByID(ctx context.Context, id int) (*model.Project, error) {
  var project *model.Project
	if err := r.Conn.QueryRowContext(
		ctx,
		"SELECT id, name, description, created_at, updated_at, completed FROM projects where id=?",
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
  conn := r.Conn
  tx, err := conn.Begin()
  if err != nil {
    return p, err
  }
  result, err := tx.ExecContext(
    ctx,
    "insert into projects (name, description, created_at, updated_at) values (?,?,?,?)",
    p.Name,
    p.Description,
    p.CreatedAt,
    p.UpdatedAt,
  )
  if err != nil {
    return p, err
  }
  id, err := result.LastInsertId()
  p.ID = int(id)
	if err = tx.Commit(); err != nil {
		return p, err
	}
  return p, nil
}



func (r *projectRepository) Update(ctx context.Context, u *model.Project) (*model.Project, error) {
}

func (r *projectRepository) Delete(ctx context.Context, id int) error {
}

// NewDB is DB constructor.
//func NewDB() (*DB, error) {
//	conn, err := sql.Open("mysql", dsn)
//	if err != nil {
//		return nil, err
//	}
//	db := DB{conn: conn}
//
//	conn.SetConnMaxLifetime(10 * time.Second)
//	conn.SetMaxOpenConns(10)
//	conn.SetMaxIdleConns(10)
//
//	if err := conn.Ping(); err != nil {
//		return nil, err
//	}
//	return &db, nil
//}
//
//// Open returns the database connection.
//func (d *DB) Open() *sql.DB {
//	return d.conn
//}
