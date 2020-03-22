package interactor

import (
	"database/sql"

	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
	"github.com/ShintaNakama/sn_project_management/app/domain/service"
	"github.com/ShintaNakama/sn_project_management/app/infrastructure/database"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
	"github.com/ShintaNakama/sn_project_management/app/usecase"
)

type Interactor interface {
	NewProjectRepository() repository.ProjectRepository
	NewProjectService() service.ProjectService
	NewProjectUseCase() usecase.ProjectUseCase
	NewProjectController() controller.ProjectController
	NewAppController() controller.AppController
}

type interactor struct {
	Conn *sql.DB
	//DB *infrastructure.DB
}

func NewInteractor(Conn *sql.DB) Interactor {
	return &interactor{Conn}
}

type appController struct {
	controller.ProjectController
}

func (i *interactor) NewAppController() controller.AppController {
	appController := &appController{}
	appController.ProjectController = i.NewProjectController()
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
