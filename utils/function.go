package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"syscall"
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
	Plural string
	Singular string
}

func GetMenu() []Menu {
	m := []Menu{
		{Url: "clients", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Plural: "Clients", Singular: "Client"},
		{Url: "modules", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail", Plural: "Modules", Singular: 	"Module"},
		{Url: "operations", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail", Plural: "Details", Singular: "Detail"	},
		{Url: "roles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Plural: "Roles", Singular: "Rol"},
		{Url: "users", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail", Change: "change", Plural: "Users", Singular: "User"},
		{Url: "dashboard", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Plural: "", Singular: ""},
		{Url: "profiles", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Change: "change", Plural: "Profiles", Singular: "Profile"},
		{Url: "contributions", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Change: "change", Plural: "Contributions", Singular: "Contribution"},
		{Url: "products", Show: "show", Create: "create",	Put: "put",	Delete: "delete",	Detail: "detail",	Change: "change", Plural: "Products", Singular: "Product"},
	}
	return m
}

func StrToTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02T15:00:00"
	t, _ := time.Parse(Format, timeStr)
	return t
}

func StrToInt(text string) (textInt int){
	integer, _ := strconv.Atoi(text)
	return integer
}

func IntToStr(num int) (txt string) {
	str := strconv.Itoa(num)
	return str
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Execute(cmd string, args string) {
	out, err := exec.Command(cmd, args).Output()

	if err != nil {
			fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

func Chdir(newdir string) {
  
  // Getting the current working directory
	CurrentWD, _ := syscall.Getwd()
	fmt.Println("CurrentWD:", CurrentWD)

	// Changing the working directory
	syscall.Chdir("/home/tech")

	// Again,
	// getting the current working directory
	CurrentWD, _ = syscall.Getwd()
	fmt.Println("CurrentWD:", CurrentWD)
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