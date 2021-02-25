package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	host_name = "192.168.1.25"
	host_port = 5432
	username = "tech"
	pwd = "F3rnand@21"
	db_name= "docusys_dev"
)


func HandleRoot(w http.ResponseWriter, r *http.Request) {

	pg_conn := fmt.Sprintf("port=%d host=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", 
	host_port, host_name, username, pwd, db_name)

	db, err := sql.Open("postgres", pg_conn)
	
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var client Client
	query := "SELECT id, first_name, last_name, ci, birthday FROM clients WHERE ci = $1"
	err = db.QueryRow(query, "3796986").Scan(&client.ID, &client.First_name, &client.Last_name, &client.Ci, &client.Birthday)
	if err != nil {
		log.Fatal("Failed to exceute query: ", err)
	}

	fmt.Fprintf(w, "Hello Wolrd!: %s %s %s %s", client.First_name, client.Last_name, client.Ci, client.Birthday)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the	API endpoint")
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