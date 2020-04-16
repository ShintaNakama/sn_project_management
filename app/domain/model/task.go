package model

import "time"

// Task is Task model
type Task struct {
	ID          int       `json:"id" db:"id, primarykey, autoincrement"`
  ProjectID   int       `json:"project_id" db:"project_id"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Completed   int       `json:"completed" db:"completed"`
}
