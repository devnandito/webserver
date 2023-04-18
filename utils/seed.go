package utils

import (
	"strings"
)


type ValidateSeed struct {
	ModuleDes string
	OperationDes string
	RoleDes string
	Errors map[string]string
}

func (msg *ValidateSeed) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.ModuleDes) == "" {
		msg.Errors["ModuleDes"] = "Please enter a module name"
	}

	if strings.TrimSpace(msg.OperationDes) == "" {
		msg.Errors["OperationDes"] = "Please enter a operation name"
	}

	if strings.TrimSpace(msg.RoleDes) == "" {
		msg.Errors["RoleDes"] = "Please enter a role name"
	}

	return len(msg.Errors) == 0

}