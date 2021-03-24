package api

import (
	"encoding/json"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var cls models.Client

// HandleApiClientGet list client
func HandleApiClientGet(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(response))
}