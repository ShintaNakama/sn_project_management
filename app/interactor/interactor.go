package interactor

import (
	"github.com/ShintaNakama/sn_project_management/app/infrastructure"
	"github.com/ShintaNakama/sn_project_management/app/presentation/http/controller"
)

type Interactor interface {
	NewProjectController() controller.ProjectController
	NewAppController() controller.AppController
}

type interactor struct {
	DB *infrastructure.DB
}

func NewInteractor(DB *infrastructure.DB) Interactor {
	return &interactor{DB}
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
	return repository.ProjectRepository(i.DB())
}

func (i *interactor) NewProjectService() service.ProjectService {
	return service.ProjectService(i.NewProjectRepository())
}

func (i *interactor) NewProjectUseCase() usecase.ProjectUseCase {
	return usecase.ProjectUseCase(i.NewProjectService())
}

func (i *interactor) NewProjectController() controller.ProjectController {
	return controller.ProjectController(i.NewProjectUseCase())
}
