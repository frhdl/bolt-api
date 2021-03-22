package taskssrv

import (
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/getchipman/bolt-api/app/core/ports"
)

// Service represent a service.
type Service struct {
	taskService ports.TaskRepository
}

// New create new instance of service.
func New(repository ports.TaskRepository) *Service {
	return &Service{
		taskService: repository,
	}
}

// Create check the parameters save a new task.
func (s *Service) Create(ctx *context.Context, task *domains.Task) context.Result {
	if task.Description == "" {
		ctx.Logger.WithField("Error", ErrorTaskDescriptionIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorTaskDescriptionIsMandatory.Error())
	}

	if task.ProjectID == 0 {
		ctx.Logger.WithField("Error", ErrorProjectIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectIDIsMandatory.Error())
	}

	return s.taskService.Create(ctx, task)
}

// GetAll check the parameters list all tasks in the target project.
func (s *Service) GetAll(ctx *context.Context, projectID int, page int, limit int) (context.Result, []*domains.Task) {

	if projectID == 0 {
		ctx.Logger.WithField("Error", ErrorProjectIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorProjectIDIsMandatory.Error()), nil
	}

	if page <= 0 {
		ctx.Logger.WithField("Error", ErrorPageMustBeGreater).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorPageMustBeGreater.Error()), nil
	}

	if limit <= 0 || limit > 100 {
		ctx.Logger.WithField("Error", ErrorPageMustBeBetween).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorPageMustBeBetween.Error()), nil
	}

	return s.taskService.GetAll(ctx, projectID, page, limit)
}

// Update check the parameters save a new data in the existing task.
func (s *Service) Update(ctx *context.Context, task *domains.Task) context.Result {
	if task.Description == "" && !task.Done {
		ctx.Logger.WithField("Error", ErrorTaskDescriptionIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorTaskDescriptionIsMandatory.Error())
	}

	if task.ID == 0 {
		ctx.Logger.WithField("Error", ErrorTaskIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorTaskIDIsMandatory.Error())
	}

	return s.taskService.Update(ctx, task)
}

// Delete check the parameters delete an existing task.
func (s *Service) Delete(ctx *context.Context, taskID int) context.Result {
	if taskID == 0 {
		ctx.Logger.WithField("Error", ErrorTaskIDIsMandatory).Errorf("Error to validate parameters")
		return ctx.ResultError(1, ErrorTaskIDIsMandatory.Error())
	}

	return s.taskService.Delete(ctx, taskID)
}
