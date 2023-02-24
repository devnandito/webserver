package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var md models.Module

// HandleApiModules list of modules
func HandleApiModules(w http.ResponseWriter, r *http.Request) {
	modules, err := md.ShowModuleGorm()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&modules)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleAPiCreateModule create new 
func HandleApiCreateModule(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&md)
	
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := md.ToJson(md)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mod := models.Module{}
	textBytes := []byte(jsonData)
	er := md.ToText(textBytes, &mod)

	if er != nil {
		panic(er)
	}

	data := models.Module{
		Description: mod.Description,
	}

	response, err := md.CreateModuleGorm(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}