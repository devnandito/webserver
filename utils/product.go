package utils

import "strings"

type ValidateProduct struct {
	Description string
	Errors map[string]string
}

func (msg *ValidateProduct) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Description) == "" {
		msg.Errors["Description"] = "Please enter a description"
	}

	return len(msg.Errors) == 0
}