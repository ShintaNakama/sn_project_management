package service

import (
	"context"

	"github.com/ShintaNakama/sn_project_management/app/domain/repository"
)

// ProjectService ドメインサービスとして利用し,複数のエンティティやレポジトリを扱う処理をここで実装する.
// ※ ドメインサービスはアプリケーションサービスではないのでトランザクションの境界などは持たない.
// ※ なんでもドメインサービスで実装するとドメインモデル貧血症となるので気をつける(ドメインモデルで表現できないかよくよく検討すること).

type ProjectService interface {
	DoSomething(ctx context.Context, foo int) error
}

type projectService struct {
	repository.ProjectRepository
}

// NewProjectService ProjectServiceを取得します.
func NewProjectService(r repository.ProjectRepository) ProjectService {
	return &projectService{r}
}

func (p *projectService) DoSomething(ctx context.Context, foo int) error {
	// some code
	return nil
}
