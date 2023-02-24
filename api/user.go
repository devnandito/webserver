package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devnandito/webserver/models"
)

var user models.User

// HandleApiUser list of user
func HandleApiUser(w http.ResponseWriter, r *http.Request) {
	users, err := user.ShowUserGorm()
	if err != nil {
		panic(err)
	}

	response, err := json.Marshal(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// HandleApiCreateUser create new user
func HandleApiCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	jsonData, err := user.ToJson(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usr := models.User{}
	textBytes := []byte(jsonData)
	er := usr.ToText(textBytes, &usr)

	if er != nil {
		panic(er)
	}

	data := models.User {
		Username: usr.Username,
		Email: usr.Email,
		Password: usr.Password,
		RoleID: usr.RoleID,
	}

	response, err := usr.CreateUserGorm(&data)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}