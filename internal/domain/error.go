package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("nothing was found")
	ErrConflict            = errors.New("this item already exists")
	ErrBadParamInput       = errors.New("given request param is not valid")
)
