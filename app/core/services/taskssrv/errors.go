package taskssrv

import "errors"

var ErrorTaskDescriptionIsMandatory = errors.New("the task 'description' field is mandatory")

var ErrorTaskIDIsMandatory = errors.New("the task 'id' field is mandatory")

var ErrorProjectIDIsMandatory = errors.New("the project 'id' field is mandatory")

var ErrorPageMustBeGreater = errors.New("page must be greater than 0")

var ErrorPageMustBeBetween = errors.New("limit must be between 1 and 100")
