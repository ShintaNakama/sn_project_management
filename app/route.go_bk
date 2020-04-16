package app

import (
	"log"
	"net/http"

	"github.com/ShintaNakama/sn_project_management/app/infrastructure"
	"github.com/ShintaNakama/sn_project_management/app/middleware"
	"github.com/ShintaNakama/sn_project_management/app/presentation"
	"github.com/gorilla/mux"
)

func NewMux() http.Handler {
	r := mux.NewRouter()

	db, err := infrastructure.NewDB()
	if err != nil {
		log.Panic(err)
	}

	//r.Use(middleware.Authenticate(db))
	r.Use(middleware.RequestLogger)
	base := presentation.Base{
		DB: db,
	}
	project := presentation.ProjectController{Base: base}
	r.HandleFunc("/projects", project.List).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id:[0-9]+}", project.Show).Methods(http.MethodGet)
	r.HandleFunc("/projects", project.Create).Methods(http.MethodPost)
	r.HandleFunc("/projects/{id:[0-9]+}", project.Update).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id:[0-9]+}", project.Delete).Methods(http.MethodDelete)

	return r
}
