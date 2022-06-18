package authsrv

import (
	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/frhdl/bolt-api/app/core/ports"
)

// Service represent a service.
type Service struct {
	authService ports.AuthRepository
}

//New create new instance of service
func New(repository ports.AuthRepository) *Service {
	return &Service{
		authService: repository,
	}
}

// Login check the login parameters.
func (s *Service) Login(ctx *context.Context, user *domains.User) (context.Result, int, string, string) {
	if user.ClientID == "" || user.ClientSecret == "" {
		ctx.Logger.WithField("Error", ErrorUserAndPasswordAreMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorUserAndPasswordAreMandatory.Error()), 0, "", ""
	}

	result, id, name, email := s.authService.Login(ctx, user)

	if result.Error == nil && id == 0 {
		ctx.Logger.WithField("Error", ErrorUserNotFound).Errorf("Error to get user")
		return ctx.ResultError(3, ErrorUserNotFound.Error()), 0, "", ""
	}

	return result, id, name, email
}
