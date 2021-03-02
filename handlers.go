package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	"log"
	"web_server/models"
	"html/template"
	// "os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wolrd!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the	API endpoint")
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	cls, err := models.AllClient()

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var listCli []Client
	for _, cl := range cls {
		data := Client{
			ID: cl.ID,
			First_name: cl.First_name,
			Last_name: cl.Last_name,
			Ci: cl.Ci,
			Birthday: cl.Birthday,
		}
		listCli = append(listCli, data)
	}
	
	tpl, err := template.ParseFiles("templates/clients/index.html")
	// t := template.Must(template.ParseFiles("templates/clients/index.html"))
	err = tpl.Execute(w, listCli)
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "applicacion/json")
	w.Write(response)
}