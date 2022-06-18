package handlers

import (
	"strconv"

	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/gin-gonic/gin"
)

// CreateTask Add a new project.
func (hdl *HTTPHandler) CreateTask(ctx *context.Context, c *gin.Context) error {
	params := &domains.Task{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.taskService.Create(ctx, params)

	result.JSON(c, nil)
	return nil
}

// GetAllTasks return all projects for logged user.
func (hdl *HTTPHandler) GetAllTasks(ctx *context.Context, c *gin.Context) error {
	pg := c.Query("page")
	lt := c.Query("limit")
	id := c.Query("projectID")

	projectID, _ := strconv.Atoi(id)

	page, err := strconv.Atoi(pg)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(lt)
	if err != nil {
		limit = 20
	}

	result, items := hdl.taskService.GetAll(ctx, projectID, page, limit)
	result.JSON(c, items)
	return nil
}

// UpdateTask Update a project.
func (hdl *HTTPHandler) UpdateTask(ctx *context.Context, c *gin.Context) error {
	params := &domains.Task{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.taskService.Update(ctx, params)

	result.JSON(c, nil)
	return nil
}

// DeleteTask Delete a project.
func (hdl *HTTPHandler) DeleteTask(ctx *context.Context, c *gin.Context) error {
	param := c.Param("taskID")
	taskID, err := strconv.Atoi(param)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.taskService.Delete(ctx, taskID)

	result.JSON(c, nil)
	return nil
}
