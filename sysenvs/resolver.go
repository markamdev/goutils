package sysenvs

import "fmt"

type VarNameResolver interface {
	SetNewPrefix(prefix string)
	GetVarName(option string) string
	GetModuleVarName(section, option string) string
}

func GetNewResolver(prefix string) VarNameResolver {
	return &vnres{
		prefixValue: prefix,
		prefixSet:   len(prefix) > 0,
	}
}

type vnres struct {
	prefixValue string
	prefixSet   bool
}

// SetNewPrefix defines prefix for all environment variables for this application
//
// Example: for Note application one can have NT_ prefix and variables like NT_PORT
// while for Alarm application it can be ALR_ prefix and ALR_PORT set in the same space
func (vr *vnres) SetNewPrefix(prefix string) {
	vr.prefixValue = prefix
}

// GetVarName returns environment variable name with prefix if any defined
//
// Example: for prefix set to GW and option name DEBUG_LEVEL variable name will be GW_DEBUG_LEVEL
func (vr *vnres) GetVarName(option string) string {
	if vr.prefixSet {
		return fmt.Sprintf("%s_%s", vr.prefixValue, option)
	}
	return option
}

// GetModuleVarName returns environment variable name with prefix (if defined) and module name included
//
// Example: for application prefix CRM and module GATEWAY the variable name for PORT will be CRM_GATEWAY_PORT
// while for the same application and option but module LOGGER it will be CRM_LOGGER_PORT
func (vr *vnres) GetModuleVarName(section, option string) string {
	if vr.prefixSet {
		return fmt.Sprintf("%s_%s_%s", vr.prefixValue, section, option)
	}
	return fmt.Sprintf("%s_%s", section, option)
}
