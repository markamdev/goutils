package sysenvs

import (
	"os"
	"strconv"
)

// GetStringEnvWithDefault return sytem env variable as string
//
// If variable with given name does not exists then empty string and
// ErrNoEnv error is returned
func GetStringEnv(name string) (string, error) {
	result := os.Getenv(name)
	if len(result) == 0 {
		return "", ErrNoEnv
	}

	return result, nil
}

// GetStringEnvWithDefault return sytem env variable as string
//
// If sytem variable with given name does not exists then defValue is returned
func GetStringEnvWithDefault(name, defValue string) string {
	result := os.Getenv(name)
	if len(result) == 0 {
		return defValue
	}

	return result
}

// GetIntEnv returns system env variable as int value
//
// If variable with given name does not exists or cant't be parsed as int
// then 0 and ErrNoEnv error is returned
func GetIntEnv(name string) (int, error) {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		return 0, ErrNoEnv
	}

	result, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, ErrInvalidEnvValue
	}

	return result, nil
}

// GetIntEnvWithDefault returns system env variable as int value
//
// If sytem variable with given name does not exist or can't be parsed as int
// then defValue is returned
func GetIntEnvWithDefault(name string, defValue int) int {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		return defValue
	}

	result, err := strconv.Atoi(strVal)
	if err != nil {
		return defValue
	}

	return result
}
