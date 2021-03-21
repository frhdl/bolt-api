package userssrv

import "errors"

var ErrorUserFieldsAreMandatory = errors.New("fields 'name', 'email', 'client_id' and 'client_secret' are mandatory")

var ErrorUserOrPasswordLegth = errors.New("fields 'client_id' or 'client_secret' must have a maximum of 16 characters each")

var ErrorUserEmailNotValid = errors.New("field 'email' is not valid")

var ErrorUserEmailAlreadyExists = errors.New("email already exists")

var ErrorUserClientIDAlreadyExists = errors.New("client_id already exists")
