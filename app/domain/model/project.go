package model

import "time"

// Project is Project model
type Project struct {
	ID          int       `json:"id" db:"id, primarykey, autoincrement"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Completed   int       `json:"completed" db:"completed"`
}
