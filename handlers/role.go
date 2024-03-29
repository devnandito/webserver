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

var rol models.Role

func HandelShowRole(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "List roles"
	headers := [3]string{"ID", "Description", "Action"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/roles", "show.html")
	response, err := rol.ShowRoleGorm()
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
		"RoleSession": roleSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template: ", res)
		return
	}
}

func HandleCreateRole(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	m := utils.GetMenu()
	title := "Add rol"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	add := filepath.Join("views/roles", "add.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[3]
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"UserSession": userSession,
			"RoleSession": roleSession,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template:", res)
			return
		}
	case "POST":
		msg := &utils.ValidateRole{
			Description: r.PostFormValue("description"),
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"UserSession": userSession,
				"RoleSession": roleSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Role {
				Description: msg.Description,
			}
			
			response, err := rol.CreateRoleGorm(&data)
	
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
				"RoleSession": roleSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		}
	}
}

func HandleUpdateRole(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Edit role"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/roles", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[3]
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		id, err :=  strconv.Atoi(sid)

		if err != nil {
			panic(err)
		}

		response, err := rol.GetOneRoleGorm(id)
		if err != nil {
			log.Println("Error executing template", response)
		}

		msg := &utils.ValidateRole {
			Description: response.Description,
		}

		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Msg": msg,
			"ID": id,
			"UserSession": userSession,
			"RoleSession": roleSession,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template", res)
			return
		}

	case "POST":
		msg := &utils.ValidateRole {
			Description: r.PostFormValue("description"),
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
				"RoleSession": roleSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Role {
				Description: msg.Description,
			}
			
			response, err := rol.UpdateRoleGorm(id, &data)
	
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
				"RoleSession": roleSession,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		}
	}
}

func HandleGetRole(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Delete role"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	delete := filepath.Join("views/roles", "delete.html")
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		panic(err)
	}
	response, err := rol.GetOneRoleGorm(id)
	if err != nil {
		panic(err)
	}

	tmpl, _ := template.ParseFiles(delete, header, nav, menu, javascript, footer)
	res := tmpl.Execute(w, map[string] interface{}{
		"Title": title,
		"Object": response,
		"UserSession": userSession,
		"RoleSession": roleSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template", res)
		return
	}
}

func HandleDeleteRole(w http.ResponseWriter, r *http.Request){
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		panic(err)
	}

	response := rol.DeleteRoleGorm(id)
	log.Println("Deleted role", response)
	http.Redirect(w, r, "/roles/show", http.StatusFound)
}