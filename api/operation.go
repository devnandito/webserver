package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var op models.Operation

// HandleApiOperation list of operation
func HandleApiOperations(w http.ResponseWriter, r *http.Request) {
	operations, err := op.ShowOperationGorm()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&operations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleApiCreateOperation
func HandleApiCreateOperation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&op)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := op.ToJson(op)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	opr := models.Operation{}
	textBytes := []byte(jsonData)
	er := op.ToText(textBytes, &opr)

	if er != nil {
		panic(er)
	}

	data := models.Operation {
		Description: opr.Description,
		ModuleID: opr.ModuleID,
	}

	response, err := op.CreateOperationGorm(&data)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}