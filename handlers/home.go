package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleHome route api
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the	API endpoint")
}

// HandlePostRequest post request
func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

// HandleUserPostRequest post user 
func HandleUserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
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
