package utils

import (
	"strings"
	"time"
)


type ValidateClient struct {
	Ci string
	Firstname string
	Lastname string
	Birthday time.Time
	Sex string
	Errors map[string]string
}

// BirthdayDateStr conver to string
func (msg ValidateClient) BirthdayDateStr() string {
	return msg.Birthday.Format("2006-01-02")
}

// BirthdayTime convert string to time
func (msg ValidateClient) BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}


func (msg *ValidateClient) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Ci) == "" {
		msg.Errors["Ci"] = "Please enter a document number"
	}

	if strings.TrimSpace(msg.Firstname) == "" {
		msg.Errors["Firstname"] = "Please enter a firstname"
	}

	if strings.TrimSpace(msg.Lastname) == "" {
		msg.Errors["Lastname"] = "Please enter a lastname"
	}

	if strings.TrimSpace(msg.Sex) == "" {
		msg.Errors["Sex"] = "Please enter a sex"
	}

	if strings.TrimSpace(msg.BirthdayDateStr()) == "" {
		msg.Errors["Birthday"] = "Please enter a birthay"
	}

	return len(msg.Errors) == 0

}

// var rxEmail = regexp.MustCompile(`.+@.+\..+`)

// type ValidateClient struct {
// 	Email string
// 	Firstname string
// 	Lastname string
// 	Errors map[string]string
// }

// func (msg *ValidateClient) Validate() bool {
// 	msg.Errors = make(map[string]string)
// 	match := rxEmail.Match([]byte(msg.Email))

// 	if !match {
// 		msg.Errors["Email"] = "Please enter a valid email address"
// 	}

// 	if strings.TrimSpace(msg.Firstname) == "" {
// 		msg.Errors["Firstname"] = "Please enter a firstname"
// 	}

// 	if strings.TrimSpace(msg.Lastname) == "" {
// 		msg.Errors["Lastname"] = "Please enter a lastname"
// 	}

// 	return len(msg.Errors) == 0

// }