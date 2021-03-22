package projectssrv

import "errors"

var ErrorProjectNameIsMandatory = errors.New("the project 'name' field is mandatory")

var ErrorProjectIDIsMandatory = errors.New("the project 'id' field is mandatory")

var ErrorPageMustBeGreater = errors.New("page must be greater than 0")

var ErrorPageMustBeBetween = errors.New("limit must be between 1 and 100")

var ErrorProjectNameAlreadyExist = errors.New("project name already exists")
