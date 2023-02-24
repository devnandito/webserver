package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var rol models.Role

// HandleApiRole list of operation
func HandleApiRole(w http.ResponseWriter, r *http.Request) {
	roles, err := rol.ShowRoleGorm()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleApiCreateRole create new role
func HandleApiCreateRole(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rol)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := rol.ToJson(rol)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rl := models.Role{}
	textBytes := []byte(jsonData)
	er := op.ToText(textBytes, &rl)

	if er != nil {
		panic(er)
	}

	data := models.Role {
		Description: rl.Description,
	}

	response, err := rol.CreateRoleGorm(&data)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}