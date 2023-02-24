package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var cls models.Client

// HandleApiClients list client
func HandleApiClients(w http.ResponseWriter, r *http.Request) {
	clients, err := cls.ShowClientGorm()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&clients)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleApiCreateClient create a new client
func HandleApiCreateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cls)
	
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := cls.ToJson(cls)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cli := models.Client{}
	textBytes := []byte(jsonData)
	er := cls.ToText(textBytes, &cli)

	if er != nil {
		panic(er)
	}

	data := models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
		Birthday: cli.Birthday,
		Sex: cli.Sex,
	}

	response, err := cls.CreateClientGorm(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonData)
}

// HandleApiCreateClient edit client
func HandleApiPutClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the	API endpoint to update")
}