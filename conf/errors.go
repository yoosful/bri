package conf

import "errors"

var (
	ErrInvalidArguments = errors.New("Invalid arguments")
	ErrDuplicateDevice  = errors.New("Duplicate Device")
)
