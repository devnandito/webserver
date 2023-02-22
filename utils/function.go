package utils

import "time"

type Link struct {
	Url string
	Action map[string]string
}

func GetUrl(url string) *Link {

	m := make(map[string]string)
	m["show"] = "show"
	m["create"] = "create"
	m["edit"] = "edit"
	m["delete"] = "delete"
	m["detail"] = "detail"
	m["success"] = "success"
	
	data := &Link{
		Url: url,
		Action: m,
	}

	return data
}

func BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}