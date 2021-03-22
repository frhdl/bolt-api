package ports

import (
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
)

// AuthRepository interface for auth service.
type AuthService interface {
	Login(*context.Context, *domains.User) (context.Result, int, string, string)
}

// UserRepository interface for user service.
type UserService interface {
	Create(*context.Context, *domains.User) context.Result
	Find(ctx *context.Context, user *domains.User) (context.Result, string, string)
}

// ProjectRepository interface for project service.
type ProjectService interface {
	Create(ctx *context.Context, project *domains.Project) context.Result
	GetAll(ctx *context.Context, page int, limit int) (context.Result, []*domains.Project)
	Update(ctx *context.Context, project *domains.Project) context.Result
	Delete(ctx *context.Context, projectID int) context.Result
	Find(ctx *context.Context, project *domains.Project) (context.Result, string)
}

// TaskRepository interface for task service.
type TaskService interface {
	Create(ctx *context.Context, task *domains.Task) context.Result
	GetAll(ctx *context.Context, projectID int, page int, limit int) (context.Result, []*domains.Task)
	Update(ctx *context.Context, task *domains.Task) context.Result
	Delete(ctx *context.Context, taskID int) context.Result
}
