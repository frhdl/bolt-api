package projectssrv

import (
	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/frhdl/bolt-api/app/core/ports"
)

// Service represent a service.
type Service struct {
	projectService ports.ProjectRepository
}

// New Create new instance of service.
func New(repository ports.ProjectRepository) *Service {
	return &Service{
		projectService: repository,
	}
}

// Create check the parameters and save a new project.
func (s *Service) Create(ctx *context.Context, project *domains.Project) context.Result {
	if project.Name == "" {
		ctx.Logger.WithField("Error", ErrorProjectNameIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectNameIsMandatory.Error())
	}

	result, name := s.projectService.Find(ctx, project)
	if result.Error != nil {
		return result
	}

	if name != "" && name == project.Name {
		ctx.Logger.WithField("Error", ErrorProjectNameAlreadyExist).Errorf("Error to validate parameters")
		return ctx.ResultError(5, ErrorProjectNameAlreadyExist.Error())
	}

	return s.projectService.Create(ctx, project)
}

// GetAll check the parameters and list all projects in the target user.
func (s *Service) GetAll(ctx *context.Context, page int, limit int) (context.Result, []*domains.Project) {
	if page <= 0 {
		ctx.Logger.WithField("Error", ErrorPageMustBeGreater).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorPageMustBeGreater.Error()), nil
	}

	if limit <= 0 || limit > 100 {
		ctx.Logger.WithField("Error", ErrorPageMustBeBetween).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorPageMustBeBetween.Error()), nil
	}

	return s.projectService.GetAll(ctx, page, limit)
}

// Update check the parameters and save a new data in the existing project.
func (s *Service) Update(ctx *context.Context, project *domains.Project) context.Result {
	if project.Name == "" {
		ctx.Logger.WithField("Error", ErrorProjectNameIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectNameIsMandatory.Error())
	}

	if project.ID == 0 {
		ctx.Logger.WithField("Error", ErrorProjectIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectIDIsMandatory.Error())
	}

	return s.projectService.Update(ctx, project)
}

// Delete check the parameters and delete an existing project.
func (s *Service) Delete(ctx *context.Context, projectID int) context.Result {
	if projectID == 0 {
		ctx.Logger.WithField("Error", ErrorProjectIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectIDIsMandatory.Error())
	}

	return s.projectService.Delete(ctx, projectID)
}

// Find check the parameters and find a projecet in database.
func (s *Service) Find(ctx *context.Context, project *domains.Project) (context.Result, string) {
	if project.ID == 0 {
		ctx.Logger.WithField("Error", ErrorProjectIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectIDIsMandatory.Error()), ""
	}

	return s.projectService.Find(ctx, project)
}
