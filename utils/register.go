package utils

import "strings"

func (msg *ValidateUser) ValidateRegister() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Username) == "" {
		msg.Errors["Username"] = "Please enter a username"
	}

	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter a name"
	}

	if strings.TrimSpace(msg.Email) == "" {
		msg.Errors["Email"] = "Please enter a email"
	}

	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please enter a password"
	}
		
	return len(msg.Errors) == 0

}