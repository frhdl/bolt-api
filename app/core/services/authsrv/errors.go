package authsrv

import "errors"

var ErrorUserAndPasswordAreMandatory = errors.New("fields 'client_id' and 'client_secret' are mandatory")

var ErrorUserNotFound = errors.New("User not found")
