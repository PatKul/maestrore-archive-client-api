package core

import "fmt"

var ErrorNotFound = fmt.Errorf("not Found")
var ErrorBadRequest = fmt.Errorf("bad Request")
var ErrorInternalServerError = fmt.Errorf("internal Server Error")
var ErrorUnauthorized = fmt.Errorf("unauthorized")
var ErrorForbidden = fmt.Errorf("forbidden")
var ErrorConflict = fmt.Errorf("conflict")
