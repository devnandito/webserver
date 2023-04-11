package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/devnandito/webserver/models"
	"github.com/devnandito/webserver/utils"
)

var cls models.Client
var user models.User
var metadata models.MetaData

func HandleShowClient(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	title := "List client"
	m := utils.GetMenu()
	headers := [7]string{"ID", "Firstname", "Lastname", "CI", "Birthday", "Sex", "Action"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/clients", "show.html")
	response, err := cls.ShowClientGorm()

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
		log.Println("Error executing template :", res)
		return
	}
}

func HandleCreateClient(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	m := utils.GetMenu()
	title := "Add client"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	add := filepath.Join("views/clients", "add.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[0]
	link := "/"+url.Url+"/"+url.Show

	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Link": link,
			"UserSession": userSession,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template:", res)
			return
		}
	case "POST":
		tBirthday := utils.BirthdayTime(r.PostFormValue("birthday"))
		msg := &utils.ValidateClient {
			Ci: r.PostFormValue("document"),
			Firstname: r.PostFormValue("firstname"),
			Lastname: r.PostFormValue("lastname"),
			Sex: r.PostFormValue("sex"),
			Birthday: tBirthday,
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"Link": link,
				"UserSession": userSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Client {
				Ci: msg.Ci,
				FirstName: msg.Firstname,
				LastName: msg.Lastname,
				Sex: msg.Sex,
				Birthday: tBirthday,
			}
			
			response, err := cls.CreateClientGorm(&data)
	
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
	
			log.Println("Data inserted", response)
			message := "Insertado correctamente"
			tmpl, _ := template.ParseFiles(ms, header, nav, menu, javascript, footer)
			linkmsg := "/"+url.Url+"/"+url.Show
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": message,
				"Link": linkmsg,
				"UserSession": userSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		}
	}
}

func HandleUpdateClient(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	title := "Edit client"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/clients", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[0]
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		id, err :=  strconv.ParseInt(sid, 10, 64)

		if err != nil {
			panic(err)
		}

		response, err := cls.GetOneClientGorm(id)
		if err != nil {
			log.Println("Error executing template", response)
		}

		msg := &utils.ValidateClient {
			Ci: response.Ci,
			Firstname: response.FirstName,
			Lastname: response.LastName,
			Sex: response.Sex,
		}

		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Msg": msg,
			"ID": id,
			"UserSession": userSession,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template", res)
			return
		}

	case "POST":
		tBirthday := utils.BirthdayTime(r.PostFormValue("birthday"))
		msg := &utils.ValidateClient {
			Ci: r.PostFormValue("document"),
			Firstname: r.PostFormValue("firstname"),
			Lastname: r.PostFormValue("lastname"),
			Sex: r.PostFormValue("sex"),
			Birthday: tBirthday,
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		
		if err != nil {
			panic(err)
		}

		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"ID": id,
				"UserSession": userSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Client {
				Ci: msg.Ci,
				FirstName: msg.Firstname,
				LastName: msg.Lastname,
				Sex: msg.Sex,
				Birthday: tBirthday,
			}
			
			response, err := cls.UpdateClientGorm(id, &data)
	
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
	
			log.Println("Data updated", response)
			message := "Actualizado correctamente"
			tmpl, _ := template.ParseFiles(ms, header, nav, menu, javascript, footer)
			linkmsg := "/"+url.Url+"/"+url.Show
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": message,
				"Link": linkmsg,
				"UserSession": userSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		}
	}
}

func HandleGetClient(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	title := "Delete client"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	delete := filepath.Join("views/clients", "delete.html")
	sid := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		panic(err)
	}
	response, err := cls.GetOneClientGorm(id)
	if err != nil {
		panic(err)
	}

	tmpl, _ := template.ParseFiles(delete, header, nav, menu, javascript, footer)
	res := tmpl.Execute(w, map[string] interface{}{
		"Title": title,
		"Object": response,
		"UserSession": userSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template", res)
		return
	}
}

func HandleDeleteClient(w http.ResponseWriter, r *http.Request){
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		panic(err)
	}

	response := cls.DeleteClientGorm(id)
	log.Println("Deleted client", response)
	http.Redirect(w, r, "/clients/show", http.StatusFound)
}

// func HandleSaveClient(w http.ResponseWriter, r *http.Request) {
// 	title := "Add Clent"
// 	tb := filepath.Join("views", "base.html")
// 	tp := filepath.Join("views/clients", "add.html")

// 	msg := &utils.Message {
// 		Email: r.PostFormValue("email"),
// 		Firstname: r.PostFormValue("firstname"),
// 		Lastname: r.PostFormValue("lastname"),
// 	}
	
// 	if !msg.Validate() {
// 		tmpl, _ := template.ParseFiles(tp, tb)
// 		res := tmpl.Execute(w, map[string]interface{}{
// 			"Title": title,
// 			"Msg": msg,
// 		})

// 		if res != nil {
// 			log.Println("Error executing template", res)
// 			return
// 		}
// 	}

// 	data := models.Client {
// 		FirstName: r.PostFormValue("firstname"),
// 		LastName: r.PostFormValue("lastname"),
// 	}
	
// 	response, err := cls.CreateClientGorm(&data)

// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	log.Println("Data inserted", response)
// 	http.Redirect(w, r, "/clients/show", http.StatusCreated)
// }

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


// HandleRoot route root
// func HandleRoot(w http.ResponseWriter, r *http.Request) {
// 		response, err := cls.ShowClientGorm()
	
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	for _, client := range response {
// 		fmt.Fprintf(w, "%d %s %s %s %s\n", client.ID, client.FirstName, client.LastName, client.Ci, client.Birthday)
// 	}
// }

// HandleShowClient list client
// func HandleShowClient(w http.ResponseWriter, r *http.Request) {
// 	response, err := cls.ShowClientGorm()

// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	tplbase := filepath.Join("templates", "layout.html")
//  	tpl := filepath.Join("templates", "show.html")
// 	res := template.Must((template.ParseFiles(tplbase, tpl)))

// 	res.Execute(w, map[string]interface{}{
// 		"Title": "List client",
// 		"clients": response,
// 	})

// 	if res != nil {
// 		log.Println("Error executing template :", res)
// 		return
// 	}
// }