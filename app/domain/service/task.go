package service

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

// TaskService ドメインサービスとして利用し,複数のエンティティやレポジトリを扱う処理をここで実装する.
// ※ ドメインサービスはアプリケーションサービスではないのでトランザクションの境界などは持たない.
// ※ なんでもドメインサービスで実装するとドメインモデル貧血症となるので気をつける(ドメインモデルで表現できないかよくよく検討すること).

type TaskService interface {
	DoSomething(ctx context.Context, foo int) error
}

type taskService struct {
	repository.TaskRepository
}

// NewProjectService ProjectServiceを取得します.
func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{r}
}

func (p *taskService) DoSomething(ctx context.Context, foo int) error {
	// some code
	return nil
}
