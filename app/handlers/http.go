package handlers

import "github.com/getchipman/bolt-api/app/core/ports"

// HTTPHandler represents a default handler.
type HTTPHandler struct {
	authService    ports.AuthService
	userService    ports.UserService
	projectService ports.ProjectService
	taskService    ports.TaskService
}

// New create a new instance of HTTPHandler.
func New(
	authService ports.AuthService,
	userService ports.UserService,
	projectService ports.ProjectService,
	taskService ports.TaskService) *HTTPHandler {
	return &HTTPHandler{
		authService:    authService,
		userService:    userService,
		projectService: projectService,
		taskService:    taskService,
	}
}
