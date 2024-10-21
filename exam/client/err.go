package client

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrNotSatisfiable = errors.New("not satisfiable")
)
