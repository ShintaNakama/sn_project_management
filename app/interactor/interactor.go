package interactor

import (
	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
	"github.com/ShintaNakama/sn_project_management/app/domain/service"
	"github.com/ShintaNakama/sn_project_management/app/infrastructure/database"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
	"github.com/ShintaNakama/sn_project_management/app/usecase"
	"github.com/go-gorp/gorp"
)

type Interactor interface {
	NewProjectRepository() repository.ProjectRepository
	NewProjectService() service.ProjectService
	NewProjectUseCase() usecase.ProjectUseCase
	NewProjectController() controller.ProjectController
	NewTaskRepository() repository.TaskRepository
	NewTaskService() service.TaskService
	NewTaskUseCase() usecase.TaskUseCase
	NewTaskController() controller.TaskController
	NewAppController() controller.AppController
}

type interactor struct {
	// database/sql
	//Conn *sql.DB
	// gorp
	Conn *gorp.DbMap
	//DB *infrastructure.DB
}

//func NewInteractor(Conn *sql.DB) Interactor {
func NewInteractor(Conn *gorp.DbMap) Interactor {
	return &interactor{Conn}
}

type appController struct {
	controller.ProjectController
	controller.TaskController
}

func (i *interactor) NewAppController() controller.AppController {
	appController := &appController{}
	appController.ProjectController = i.NewProjectController()
	appController.TaskController = i.NewTaskController()
	return appController
}

func (i *interactor) NewProjectRepository() repository.ProjectRepository {
	return database.NewProjectRepository(i.Conn)
}

func (i *interactor) NewProjectService() service.ProjectService {
	return service.NewProjectService(i.NewProjectRepository())
}

func (i *interactor) NewProjectUseCase() usecase.ProjectUseCase {
	return usecase.NewProjectUseCase(i.NewProjectRepository())
}

func (i *interactor) NewProjectController() controller.ProjectController {
	return controller.NewProjectController(i.NewProjectUseCase())
}

func (i *interactor) NewTaskRepository() repository.TaskRepository {
	return database.NewTaskRepository(i.Conn)
}

func (i *interactor) NewTaskService() service.TaskService {
	return service.NewTaskService(i.NewTaskRepository())
}

func (i *interactor) NewTaskUseCase() usecase.TaskUseCase {
	return usecase.NewTaskUseCase(i.NewTaskRepository())
}

func (i *interactor) NewTaskController() controller.TaskController {
	return controller.NewTaskController(i.NewTaskUseCase())
}
