package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/devnandito/webserver/models"
)

// TemplateRegistry initial
// type TemplateRegistry struct {
// 	templates map[string]*template.Template
// }

// func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}) error {
// 	tmpl, ok := t.templates[name]
// 	if !ok {
// 		err := errors.New("Template not found"+ name)
// 		return err
// 	}
// 	return tmpl.ExecuteTemplate(w, "base.html", data)
// }


var cls models.Client
var user models.User
var metadata models.MetaData

// HandleRoot route root
func HandleRoot(w http.ResponseWriter, r *http.Request) {
		response, err := cls.ShowClientGorm()
	
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, client := range response {
		fmt.Fprintf(w, "%d %s %s %s %s\n", client.ID, client.FirstName, client.LastName, client.Ci, client.Birthday)
	}
}

// HandleShowClient list client
func HandleShowClient(w http.ResponseWriter, r *http.Request) {
	response, err := cls.ShowClientGorm()

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	parsedTemplate, _ := template.ParseFiles("views/clients/index.html", "views/base.html")
	res := parsedTemplate.Execute(w, map[string]interface{}{
		"Title": "List client",
		"clients": response,
	})

	if res != nil {
		log.Println("Error executing template :", res)
		return
	}
}
