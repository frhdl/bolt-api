package ports

import (
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
)

// AuthRepository interface for auth repository.
type AuthRepository interface {
	Login(*context.Context, *domains.User) (context.Result, int, string, string)
}

// UserRepository interface for user repository.
type UserRepository interface {
	Create(*context.Context, *domains.User) context.Result
	Find(ctx *context.Context, user *domains.User) (context.Result, string, string)
}

// ProjectRepository interface for project repository.
type ProjectRepository interface {
	Create(ctx *context.Context, project *domains.Project) context.Result
	GetAll(ctx *context.Context, page int, limit int) (context.Result, []*domains.Project)
	Update(ctx *context.Context, project *domains.Project) context.Result
	Delete(ctx *context.Context, projectID int) context.Result
	Find(ctx *context.Context, project *domains.Project) (context.Result, string)
}

// TaskRepository interface for task repository.
type TaskRepository interface {
	Create(ctx *context.Context, task *domains.Task) context.Result
	GetAll(ctx *context.Context, projectID int, page int, limit int) (context.Result, []*domains.Task)
	Update(ctx *context.Context, task *domains.Task) context.Result
	Delete(ctx *context.Context, taskID int) context.Result
}
