package userssrv

import (
	"github.com/getchipman/bolt-api/app/common"
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/getchipman/bolt-api/app/core/ports"
)

// Service represent a service.
type Service struct {
	userService ports.UserRepository
}

//New create new instance of service.
func New(repository ports.UserRepository) *Service {
	return &Service{
		userService: repository,
	}
}

// Create check the parameters to save a new user.
func (s *Service) Create(ctx *context.Context, user *domains.User) context.Result {

	if user.Name == "" || user.Email == "" || user.ClientID == "" || user.ClientSecret == "" {
		ctx.Logger.WithField("Error", ErrorUserFieldsAreMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorUserFieldsAreMandatory.Error())
	}

	if len(user.ClientID) > 16 || len(user.ClientSecret) > 16 {
		ctx.Logger.WithField("Error", ErrorUserOrPasswordLegth).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorUserOrPasswordLegth.Error())
	}

	if !common.IsEmailValid(user.Email) {
		ctx.Logger.WithField("Error", ErrorUserEmailNotValid).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorUserEmailNotValid.Error())
	}

	// checks if the user exists
	result, email, clientID := s.userService.Find(ctx, user)
	if result.Error != nil {
		return result
	}

	if email != "" && email == user.Email {
		ctx.Logger.WithField("Error", ErrorUserEmailAlreadyExists).Errorf("Error to validate parameters")
		return ctx.ResultError(5, ErrorUserEmailAlreadyExists.Error())
	}

	if clientID != "" && clientID == user.ClientID {
		ctx.Logger.WithField("Error", ErrorUserClientIDAlreadyExists).Errorf("Error to validate parameters")
		return ctx.ResultError(5, ErrorUserClientIDAlreadyExists.Error())
	}

	return s.userService.Create(ctx, user)
}

// Find check the parameters and find a user.
func (s *Service) Find(ctx *context.Context, user *domains.User) (context.Result, string, string) {
	if user.Name == "" || user.Email == "" {
		ctx.Logger.WithField("Error", ErrorUserFieldsAreMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorUserFieldsAreMandatory.Error()), "", ""
	}

	return s.userService.Find(ctx, user)
}
