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
