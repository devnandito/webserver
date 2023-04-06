package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devnandito/webserver/models"
	"github.com/devnandito/webserver/utils"
)

var usr models.User

func HandelShowUser(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	title := "List users"
	headers := [5]string{"ID", "Username", "Email", "Full name", "Action"}
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
		"Users": response,
		"Headers": headers,
		"User": userSession,
		"Menu": m,
	})

	if res != nil {
		log.Println("Error executing template: ", res)
		return
	}
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
			Fullname: r.PostFormValue("fullname"),
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
				Name: msg.Fullname,
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