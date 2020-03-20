package router

import (
	"net/http"

	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/middleware"
	"github.com/gorilla/mux"
)

func NewMux() http.Handler {
	r := mux.NewRouter()

	//db, err := infrastructure.NewDB()
	//if err != nil {
	//	log.Panic(err)
	//}

	//r.Use(middleware.Authenticate(db))
	r.Use(middleware.RequestLogger)
	//base := controller.Base{
	//	DB: db,
	//}
	//project := controller.ProjectController{Base: base}
	r.HandleFunc("/projects", controller.AppController.List).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id:[0-9]+}", controller.AppController.Show).Methods(http.MethodGet)
	r.HandleFunc("/projects", controller.AppController.Create).Methods(http.MethodPost)
	r.HandleFunc("/projects/{id:[0-9]+}", controller.AppController.Update).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id:[0-9]+}", controller.AppController.Delete).Methods(http.MethodDelete)

	return r
}
