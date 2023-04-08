package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devnandito/webserver/models"
	"github.com/devnandito/webserver/utils"
)

var rol models.Role

func HandelShowRole(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	title := "List roles"
	headers := [3]string{"ID", "Description", "Action"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/roles", "show.html")
	response, err := rol.ShowRoleGorm()
	userSession := session.Values["username"]

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	tmpl, _ := template.ParseFiles(show, header, nav, menu, javascript, footer)
	res := tmpl.Execute(w, map[string]interface{}{
		"Title": title,
		"Objects": response,
		"Headers": headers,
		"UserSession": userSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template: ", res)
		return
	}
}