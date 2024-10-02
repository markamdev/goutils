package sysenvs

import (
	"fmt"
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

// GetMandatoryStringEnv works as GetStringEnv but panic instead of returning error
func GetMandatoryStringEnv(name string) string {
	result := os.Getenv(name)
	if len(result) == 0 {
		panic(fmt.Sprintf("cannot get %s variable and return as string", name))
	}

	return result
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
// then 0 and appropriate error is returned
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

// GetMandatoryIntEnv works as GetIntEnv but panic instead of returning error
func GetMandatoryIntEnv(name string) int {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		panic(fmt.Sprintf("cannot get %s variable and return as int", name))
	}

	result, err := strconv.Atoi(strVal)
	if err != nil {
		panic(fmt.Sprintf("cannot get %s variable and return as int", name))
	}

	return result
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

// GetBoolEnv returns system env variable as bool value
//
// If variable with given name does not exists or cant't be parsed as bool
// then false and appropriate error is returned
//
// Accepted env values are:
// TRUE, true or 1 for bool true
// FALSE, false or 0 for bool false
func GetBoolEnv(name string) (bool, error) {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		return false, ErrNoEnv
	}

	switch strVal {
	case "TRUE", "true", "1":
		return true, nil
	case "FALSE", "false", "0":
		return false, nil
	default:
		return false, ErrInvalidEnvValue
	}
}

// GetMandatoryBoolEnv as GetBoolEnv but panic instead of returning error
func GetMandatoryBoolEnv(name string) bool {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		panic(fmt.Sprintf("cannot get %s variable and return as bool", name))
	}

	switch strVal {
	case "TRUE", "true", "1":
		return true
	case "FALSE", "false", "0":
		return false
	default:
		panic(fmt.Sprintf("cannot get %s variable and return as bool", name))
	}
}

// GetBoolEnvWithDefault returns system env variable as bool value
//
// If sytem variable with given name does not exist or can't be parsed as bool
// then defValue is returned
//
// Accepted env values are:
// TRUE, true or 1 for bool true
// FALSE, false or 0 for bool false
func GetBoolEnvWithDefault(name string, defValue bool) bool {
	strVal := os.Getenv(name)
	if len(strVal) == 0 {
		return defValue
	}

	switch strVal {
	case "TRUE", "true", "1":
		return true
	case "FALSE", "false", "0":
		return false
	default:
		return defValue
	}
}
