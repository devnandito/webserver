package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devnandito/webserver/models"
	"github.com/devnandito/webserver/utils"
)

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
		
		if !msg.ValidateRegister() {
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