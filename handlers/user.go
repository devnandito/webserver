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

func HandelShowUser(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	title := "List users"
	headers := [6]string{"ID", "Username", "Email", "Full name", "Role", "Action"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/users", "show.html")
	response, err := usr.ShowUserGorm()
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

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	m := utils.GetMenu()
	title := "Add user"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	add := filepath.Join("views/users", "add.html")
	ms := filepath.Join("views/messages", "message.html")
	url := utils.GetUrl("users")
	link := "/"+url.Url+"/"+url.Action["create"]
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
		roleId, _ := strconv.Atoi(r.PostFormValue("role"))
		msg := &utils.ValidateUser{
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
			Username: r.PostFormValue("username"),
			Name: r.PostFormValue("name"),
			RoleID: roleId,

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
			data := models.User {
				Username: msg.Username,
				Email: msg.Email,
				Name: msg.Name,
				RoleID: int(msg.RoleID),
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
			linkmsg := "/"+url.Url+"/"+url.Action["show"]
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

func HandleUpdateUser(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	title := "Edit User"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/users", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := utils.GetUrl("users")
	link := "/"+url.Url+"/"+url.Action["edit"]
	
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
			RoleID: response.RoleID,
		}

		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Msg": msg,
			"Link": link,
			"ID": id,
			"UserSession": userSession,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template", res)
			return
		}

	case "POST":
		roleId, _ := strconv.Atoi(r.PostFormValue("role"))
		msg := &utils.ValidateUser {
			Username: r.PostFormValue("username"),
			Email: r.PostFormValue("email"),
			Name: r.PostFormValue("name"),
			RoleID: roleId,
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		
		if err != nil {
			panic(err)
		}

		if !msg.ValidateEdit() {
			tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"Link": link,
				"ID": id,
				"UserSession": userSession,
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
				RoleID: msg.RoleID,
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
			linkmsg := "/"+url.Url+"/"+url.Action["show"]
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": message,
				"Link": linkmsg,
				"User": userSession,
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
	title := "Delete user"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	delete := filepath.Join("views/users", "delete.html")
	sid := r.URL.Query().Get("id")
	url := utils.GetUrl("users")
	link := "/"+url.Url+"/"+url.Action["delete"]+"?id="+sid
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
		"User": response,
		"Link": link,
		"UserSession": userSession,
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
	switch r.Method {
	case "GET":
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
			session.Save(r, w)
			http.Redirect(w, r, "/users/show", http.StatusFound)
		}
	}
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	title := "Register"
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	signup := filepath.Join("views/users", "signup.html")
	ms := filepath.Join("views/messages", "msg.html")
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(signup, header, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
		})

		if res != nil {
			log.Println("Error executing template: ", res)
			return
		}
	case "POST":
		msg := &utils.ValidateUser{
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
			Username: r.PostFormValue("username"),
			Name: r.PostFormValue("name"),
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(signup, header, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}

		} else {
			pwd, _ := usr.GetPwdHash(msg.Password)
			data := models.User {
				Email: msg.Email,
				Username: msg.Username,
				Name: msg.Name,
				RoleID: 1,
				Password: pwd,
			}

			response, err := usr.CreateUserGorm(&data)

			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}

			log.Println("Data inserted", response)
			message := "Usuario registrado"
			tmpl, _ := template.ParseFiles(ms, header, footer)
			linkmsg := "/"
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": message,
				"Link": linkmsg,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}
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