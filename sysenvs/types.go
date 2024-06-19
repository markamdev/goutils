package sysenvs

import "errors"

var (
	ErrNoEnv           = errors.New("no env variable found")
	ErrInvalidEnvValue = errors.New("environment variable with invalid value/format")
)
