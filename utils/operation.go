package utils

import (
	"strings"
)


type ValidateOperation struct {
	Description string
	ModuleID string
	Errors map[string]string
}


func (msg *ValidateOperation) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Description) == "" {
		msg.Errors["Description"] = "Please enter a description"
	}

	if strings.TrimSpace(msg.ModuleID) == "" {
		msg.Errors["ModuleID"] = "Please enter a module"
	}

	return len(msg.Errors) == 0

}