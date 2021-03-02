package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Middleware func (http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Client struct {
	ID int
	First_name string
	Last_name string
	Ci string
	Birthday time.Time
}

type Table4 struct {
	PK int
	Email string
	Fullname string
}


func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}