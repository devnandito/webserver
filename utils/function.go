package utils

import "time"

type Link struct {
	Url string
	Action map[string]string
}

type Menu struct {
	Url string
	Action map[int]string
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

func GetMenu() map[string][]string {
	clients := []string{"clients", "show", "create", "edit", "detail", "delete"}
	users := []string{"users", "show", "create", "edit", "detail", "delete"}
	modules := []string{"modules", "show", "create", "edit", "detail", "delete"}
	operations := []string{"operations", "show", "create", "edit", "detail", "delete"}
	roles := []string{"roles", "show", "create", "edit", "detail", "delete"}
	m := map[string][]string{"clients": clients, "users": users, "modules": modules, "operations": operations, "roles": roles}
	return m
}

func BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}