package controller

import "github.com/ShintaNakama/sn_project_management/app/infrastructure"

// Base abstract controllr
type Base struct {
	DB *infrastructure.DB
}
