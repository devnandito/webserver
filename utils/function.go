package utils

import (
	"time"
)

type Link struct {
	Url string
	Action map[string]string
}

type Menu struct {
	Url string
	Show string
	Create string
	Put string
	Delete string
	Detail string
	Change string
}

func GetMenu() []Menu {
	m := []Menu{
		{Url: "clients", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "modules", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "operations", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "roles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "users", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail", Change: "change"},
		{Url: "dashboard", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	},
		{Url: "profiles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Change: "change"},
	}
	return m
}

func BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:04:05"
	t, _ := time.Parse(Format, timeStr)
	return t
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// type FormModule struct {
// 	Value string
// 	Option string
// 	Selected string
// }

// func SelectModuleOption(pk int, objects []models.Module) []FormModule {
// 	for _, v := range objects {
// 		if int(v.ID) == pk {
// 			rs = append(rs, FormModule{
// 				Value: strconv.FormatUint(uint64(v.ID), 10),
// 				Option: v.Description,
// 				Selected: "selected",
// 			})
// 		} else if int(v.ID) != pk {
// 			rs = append(rs, FormModule{
// 				Value: strconv.FormatUint(uint64(v.ID), 10),
// 				Option: v.Description,
// 			})
// 		}
// 	}
// 	return rs
// }

// func GetUrl(url string) *Link {
// 	m := make(map[string]string)
// 	m["show"] = "show"
// 	m["create"] = "create"
// 	m["put"] = "put"
// 	m["delete"] = "delete"
// 	m["detail"] = "detail"
// 	m["success"] = "success"
	
// 	data := &Link{
// 		Url: url,
// 		Action: m,
// 	}

// 	return data
// }