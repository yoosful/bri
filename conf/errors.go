// Copyright © 2018 Seonghyun Park <pseohy@gmail.com>

package conf

import "errors"

var (
	ErrInvalidArguments = errors.New("Invalid arguments")
	ErrDuplicateDevice  = errors.New("Duplicate device")
	ErrDuplicateUser    = errors.New("Duplicate user")
	ErrNoMathingDevice  = errors.New("No matching device")
	ErrNoMathingUser    = errors.New("No matching user")
)
