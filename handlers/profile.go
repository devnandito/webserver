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

func HandleProfileUser(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Profile User"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	profile := filepath.Join("views/profiles", "show.html")
	headers := [6]string{"ID", "Username", "Email", "Full name", "Role", "Action"}
	tmpl, _ := template.ParseFiles(profile, header, nav, menu, javascript, footer)
	username := r.URL.Query().Get("username")
	object, err := usr.SearchUsername(username)
	if err != nil {
		log.Println(err)
	}

	res := tmpl.Execute(w, map[string]interface{}{
		"Title": title,
		"Headers": headers,
		"Object": object,
		"UserSession": userSession,
		"RoleSession": roleSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template", res)
		return
	}
}

func HandleUpdateProfile(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Edit User"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/profiles", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[6]
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		id, err :=  strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		response, err := usr.GetOneUserGorm(id)
		if err != nil {
			log.Println("Error executing template", response)
		}

		msg := &utils.ValidateUser {
			Username: response.Username,
			Email: response.Email,
			Name: response.Name,
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
		msg := &utils.ValidateUser {
			Username: r.PostFormValue("username"),
			Email: r.PostFormValue("email"),
			Name: r.PostFormValue("name"),
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		if !msg.ValidateProfile() {
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
			data := models.User {
				Username: msg.Username,
				Email: msg.Email,
				Name: msg.Name,
			}
			
			response, err := usr.UpdateUserGorm(id, &data)
	
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
	
			log.Println("Data updated", response)
			message := "Actualizado correctamente"
			tmpl, _ := template.ParseFiles(ms, header, nav, menu, javascript, footer)
			linkmsg := "/"+url.Url+"/"+url.Show+"?username="+msg.Username
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

func HandleChangeProfilePwd(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Change password"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/profiles", "changepwd.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[6]
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		id, err :=  strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		response, err := usr.GetOneUserGorm(id)
		if err != nil {
			log.Println("Error executing template", response)
		}


		msg := &utils.ValidateUser {
			Password: response.Password,
			Password1: response.Password,
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
		msg := &utils.ValidateUser {
			Password: r.PostFormValue("password"),
			Password1: r.PostFormValue("password1"),
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		if !msg.ValidatePwdProfile() {
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
			pwd, _ := usr.GetPwdHash(msg.Password)
			data := models.User {
				Password: pwd,
			}
			
			response, err := usr.UpdatePwdUserGorm(id, &data)
	
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
	
			log.Println("Data updated", response)

			us, err := usr.GetOneUserGorm(id)
			if err != nil {
				log.Println(err)
			}

			message := "Actualizado correctamente"
			tmpl, _ := template.ParseFiles(ms, header, nav, menu, javascript, footer)
			linkmsg := "/"+url.Url+"/"+url.Show+"?username="+us.Username
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