package controller

// AppController interfase
// AppControllerは全てのHandlerのinterfaceを満たす.※routerの実装が依存する.
type AppController interface {
  ProjectController
    // embed all handler interfaces
}
