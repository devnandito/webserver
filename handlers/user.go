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

var usr models.User

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Dashboard"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/users", "dashboard.html")

	tmpl, _ := template.ParseFiles(show, header, nav, menu, javascript, footer)
	res := tmpl.Execute(w, map[string]interface{}{
		"Title": title,
		"UserSession": userSession,
		"RoleSession": roleSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template: ", res)
		return
	}
}

func HandleShowUser(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "List users"
	headers := [6]string{"ID", "Username", "Email", "Full name", "Role", "Action"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/users", "show.html")
	response, err := usr.ShowUserGorm()
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

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	m := utils.GetMenu()
	title := "Add user"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	add := filepath.Join("views/users", "add.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[4]

	roles, err := rol.ShowRoleGorm()
	if err != nil {
		log.Println(err)
	}

	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"UserSession": userSession,
			"RoleSession": roleSession,
			"Roles": roles,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template:", res)
			return
		}
	case "POST":
		msg := &utils.ValidateUser{
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
			Username: r.PostFormValue("username"),
			Name: r.PostFormValue("name"),
			RoleID: r.PostFormValue("role"),
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"UserSession": userSession,
				"Roles": roles,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			id, err := strconv.Atoi(msg.RoleID)
			if err != nil {
				log.Println(err)
			}
			pwd, _ := usr.GetPwdHash(msg.Password)

			data := models.User {
				Username: msg.Username,
				Email: msg.Email,
				Name: msg.Name,
				RoleID: id,
				Password: pwd,
			}
			
			response, err := usr.CreateUserGorm(&data)
	
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

func HandleUpdateUser(w http.ResponseWriter, r *http.Request){
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
	edit := filepath.Join("views/users", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[4]
	roles, err := rol.ShowRoleGorm()
	if err != nil {
		log.Println(err)
	}
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		// id, err :=  strconv.ParseInt(sid, 10, 64)
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
			RoleID: strconv.Itoa(response.RoleID),
		}

		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Msg": msg,
			"ID": id,
			"UserSession": userSession,
			"RoleSession": roleSession,
			"Roles": roles,
			"FK": response.RoleID,
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
			RoleID: r.PostFormValue("role"),
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		rolid, err := strconv.Atoi(r.PostFormValue("role"))
		if err != nil {
			log.Println(err)
		}

		if !msg.ValidateEdit() {
			tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"ID": id,
				"UserSession": userSession,
				"RoleSession": roleSession,
				"Roles": roles,
				"FK": rolid,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			rolid, err := strconv.Atoi(msg.RoleID)
			if err != nil {
				log.Println(err)
			}

			data := models.User {
				Username: msg.Username,
				Email: msg.Email,
				Name: msg.Name,
				RoleID: rolid,
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

func HandleChangePwd(w http.ResponseWriter, r *http.Request){
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
	edit := filepath.Join("views/users", "changepwd.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[4]
	
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
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		if err != nil {
			panic(err)
		}

		if !msg.ValidatePwd() {
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

func HandleGetUser(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Delete user"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	delete := filepath.Join("views/users", "delete.html")
	sid := r.URL.Query().Get("id")
	//id, err := strconv.ParseInt(sid, 10, 64)
	id, err := strconv.Atoi(sid)
	if err != nil {
		panic(err)
	}
	response, err := usr.GetOneUserGorm(id)
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

func HandleDeleteUser(w http.ResponseWriter, r *http.Request){
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		panic(err)
	}

	response := usr.DeleteUserGorm(id)
	log.Println("Deleted user", response)
	http.Redirect(w, r, "/users/show", http.StatusFound)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	title := "Login"
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	signin := filepath.Join("views/users", "signin.html")
	link := "/login"
	session := utils.GetSession(r)
	switch r.Method {
	case "GET":
		if session.Values["authenticated"] == true {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
		}
		tmpl, _ := template.ParseFiles(signin, header, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Link": link,
		})

		if res != nil {
			log.Println("Error executing template: ", res)
			return
		}
	case "POST":
		msg := &utils.ValidateUser {
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		
		if !msg.ValidateLogin() {
			tmpl, _ := template.ParseFiles(signin, header, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"Link": link,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}
		} else {
			session := utils.GetSession(r)
			userSession := utils.GetUser(r.PostFormValue("email"))
			session.Values["authenticated"] = true
			session.Values["username"] = userSession.Username
			session.Values["role"] = userSession.RoleID
			session.Save(r, w)
			http.Redirect(w, r, "/dashboard", http.StatusFound)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}