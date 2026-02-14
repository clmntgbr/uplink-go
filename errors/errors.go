package errors

import "errors"

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrUserNotFound    = errors.New("user not found")
)