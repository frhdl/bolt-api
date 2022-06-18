package handlers

import (
	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/gin-gonic/gin"
)

// Create Add a new user.
func (hdl *HTTPHandler) Create(ctx *context.Context, c *gin.Context) error {

	params := &domains.User{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
	}

	result := hdl.userService.Create(ctx, params)

	result.JSON(c, nil)
	return nil
}
