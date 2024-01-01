package common

import "errors"

var (
	RowNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)
