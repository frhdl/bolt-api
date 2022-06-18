package context

import (
	"fmt"

	"github.com/frhdl/bolt-api/app/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	environmentdev  = "dev"
	environmentprod = "prod"
)

// Context is a struct that contains a global fields for context.
type Context struct {
	UUID         string `json:"uuid"`
	LoggedUserID int
	IsAdmin      bool
	HTTPPrefix   string
	Logger       *log.Entry
}

// ResultState represent actual state of result API response.
type ResultState int

// Result represent API result.
type Result struct {
	State            ResultState
	Error            error
	ErrorCode        int
	ErrorDescription string
}

const (
	// ResultStateSuccess state is when the result is successful.
	ResultStateSuccess ResultState = 0

	// ResultStateError state is when the result has validation error.
	ResultStateError ResultState = 1

	// ResultStateUnexpected state is when the result has unexpected error.
	ResultStateUnexpected ResultState = 2

	// ResultStateNotFound state is when a register was not found.
	ResultStateNotFound ResultState = 3

	// ResultStateUnauthorized state is when the request is not authorized.
	ResultStateUnauthorized ResultState = 4

	// ResultStateResourceAlreadyExist state is when the resource already exist.
	ResultStateResourceAlreadyExist = 5
)

// ResultError represent a base for applications errors.
type ResultError struct {
	Code        int    `json:"error_code"`
	Description string `json:"error_description"`
}

// New create a new context instance with UUID.
func New() *Context {
	return &Context{
		UUID: common.GenerateUUID(),
	}
}

// WithLogger add logger to context for better logs.
func (c *Context) WithLogger() *Context {
	contextLogger := log.WithFields(log.Fields{
		"uuid": c.UUID,
	})

	if common.GetEnv("environment", environmentdev) == environmentprod {
		log.SetReportCaller(true)
		log.SetFormatter(&log.JSONFormatter{})
	}

	c.Logger = contextLogger
	return c
}

// ResultSuccess build a common struct for success.
func (c *Context) ResultSuccess() Result {
	result := Result{
		State: ResultStateSuccess,
	}

	return result
}

// ResultError build a common struct for validation error result.
func (c *Context) ResultError(errorCode int, description string, params ...interface{}) Result {

	if len(params) > 0 {
		description = fmt.Sprintf(description, params...)
	}

	result := Result{
		State:            ResultStateError,
		ErrorCode:        errorCode,
		ErrorDescription: description,
	}

	c.Logger.Warningf("Code: %v - Message: %v\n", errorCode, description)

	return result
}

// ResultNotFound build a common struct for notfound record result.
func (c *Context) ResultNotFound(errorCode int, description string, params ...interface{}) Result {

	if len(params) > 0 {
		description = fmt.Sprintf(description, params...)
	}

	result := Result{
		State:            ResultStateNotFound,
		ErrorCode:        errorCode,
		ErrorDescription: description,
	}

	c.Logger.Warningf("Code: %v - Message: %v\n", errorCode, description)

	return result
}

// ResultUnauthorized build a common struct for unauthorized requests.
func (c *Context) ResultUnauthorized() Result {
	result := Result{
		State: ResultStateUnauthorized,
	}

	return result
}

// ResultUnexpected build a common struct for unexpected error result.
func (c *Context) ResultUnexpected(err error) Result {
	errorReturn := 999999
	description := "You shouldn't get this error. But probably will be fixed in the next version."

	result := Result{
		State:            ResultStateUnexpected,
		Error:            err,
		ErrorCode:        errorReturn,
		ErrorDescription: description,
	}

	c.Logger.Error(description)

	return result
}

// JSON convert Result to HTTP Response.
func (r Result) JSON(c *gin.Context, body interface{}) {
	if r.State == ResultStateSuccess {
		if body == nil {
			c.String(200, "")
		} else {
			c.JSON(200, body)
		}
	} else {
		errorReturn := ResultError{
			Description: r.ErrorDescription,
			Code:        r.ErrorCode,
		}

		if r.State == ResultStateError {
			c.JSON(400, errorReturn)
		} else if r.State == ResultStateNotFound {
			c.JSON(404, errorReturn)
		} else if r.State == ResultStateUnauthorized {
			c.String(401, "")
		} else {
			c.JSON(500, errorReturn)
		}
	}
}
