package denv

import "errors"

var (
	// ErrEnvParse is returned when the environment variable cannot be parsed.
	ErrEnvParse = errors.New("cannot parse environment variable")
	// ErrInvalidEnvironment is returned when the environment is not one of the allowed values.
	ErrInvalidEnvironment = errors.New("invalid environment")
)
