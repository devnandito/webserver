package utils

import (
	"strconv"
	"strings"
	"time"
)

type ValidateContribution struct {
	Amount int
	Contribution_date time.Time
	Method_pay string
	Description string
	ClientID int
	ProductID int
	Errors map[string]string
}

func (msg ValidateContribution) ContDateStr() string {
	return msg.Contribution_date.Format("2006-01-02T15:04:05")
}

func (msg ValidateContribution) AmountStr() (txt string) {
	amountStr := strconv.Itoa(msg.Amount)
	return amountStr
}

func (msg ValidateContribution) StrToDate(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}

func (msg *ValidateContribution) Validate() bool {
	msg.Errors = make(map[string]string)

	if msg.Amount == 0 {
		msg.Errors["Amount"] = "Please enter a amount"
	}

	// if msg.Contribution_date.IsZero() {
	// 	msg.Errors["Contribution_date"] = "Pleasea enter a date"
	// }

	if strings.TrimSpace(msg.Method_pay) == "" {
		msg.Errors["Method_pay"] = "Please enter a method pay"
	}

	if strings.TrimSpace(msg.Description) == "" {
		msg.Errors["Description"] = "Please enter a description"
	}

	if msg.ClientID == 0 {
		msg.Errors["ClientID"] = "Please enter a client"
	}

	if msg.ProductID == 0 {
		msg.Errors["ProductID"] = "Please enter a product"
	}

	return len(msg.Errors) == 0
}