package handlers

import (
	"strconv"

	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/gin-gonic/gin"
)

// CreateProject Add a new project.
func (hdl *HTTPHandler) CreateProject(ctx *context.Context, c *gin.Context) error {
	params := &domains.Project{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.projectService.Create(ctx, params)

	result.JSON(c, nil)
	return nil
}

// GetAllProjects return all projects for logged user.
func (hdl *HTTPHandler) GetAllProjects(ctx *context.Context, c *gin.Context) error {
	pg := c.Query("page")
	lt := c.Query("limit")

	page, err := strconv.Atoi(pg)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(lt)
	if err != nil {
		limit = 20
	}

	result, items := hdl.projectService.GetAll(ctx, page, limit)
	result.JSON(c, items)
	return nil
}

// UpdateProject Update a project.
func (hdl *HTTPHandler) UpdateProject(ctx *context.Context, c *gin.Context) error {
	params := &domains.Project{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.projectService.Update(ctx, params)

	result.JSON(c, nil)
	return nil
}

// DeleteProject Delete a project.
func (hdl *HTTPHandler) DeleteProject(ctx *context.Context, c *gin.Context) error {
	param := c.Param("projectID")
	projectID, err := strconv.Atoi(param)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.projectService.Delete(ctx, projectID)

	result.JSON(c, nil)
	return nil
}
