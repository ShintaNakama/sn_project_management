package controller

import (
	"encoding/json"
	"net/http"
)

type ProjectController interface {
  List(w http.ResponseWriter, r *http.Request)
  Show(w http.ResponseWriter, r *http.Request)
  Create(w http.ResponseWriter, r *http.Request)
  Update(w http.ResponseWriter, r *http.Request)
  Delete(w http.ResponseWriter, r *http.Request)
}
type projectController struct {
	//Base
	//mux sync.Mutex
  // usecase を持ってくる
  //ProjectUseCase usecase.ProjectUseCase
}

//func NewProjectController(u usecase.ProjectUseCase){
//  return &projectController{u}
//}

// List is Get Projects
func (p *projectController) List(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	list := "project list"
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Show is Get Project
func (p *projectController) Show(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	list := "project show"
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create is Create Project
func (p *projectController) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	list := "project create"
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update is Update Project
func (p *projectController) Update(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	list := "project update"
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete is Delete Project
func (p *projectController) Delete(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	list := "project delete"
	if err := json.NewEncoder(w).Encode(list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
