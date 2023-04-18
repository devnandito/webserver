package utils

import (
	"strings"
)


type ValidateInstall struct {
	UserDB string
	PwdDB string
	NameDB string
	PortDB string
	HostDB string
	Errors map[string]string
}

func (msg *ValidateInstall) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.UserDB) == "" {
		msg.Errors["UserDB"] = "Please enter a database user"
	}

	if strings.TrimSpace(msg.PwdDB) == "" {
		msg.Errors["PwdDB"] = "Please enter a database password"
	}

	if strings.TrimSpace(msg.NameDB) == "" {
		msg.Errors["NameDB"] = "Please enter a database name"
	}

	if strings.TrimSpace(msg.HostDB) == "" {
		msg.Errors["HostDB"] = "Please enter a database host"
	}

	return len(msg.Errors) == 0

}