package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/devnandito/webserver/utils"
)

func HandleInstall(w http.ResponseWriter, r *http.Request) {
	title := "Install"
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	install := filepath.Join("views/install", "install.html")
	ms := filepath.Join("views/messages", "msg.html")
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(install, header, footer)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
		})

		if res != nil {
			log.Println("Error executing template: ", res)
			return
		}
	case "POST":
		msg := &utils.ValidateInstall{
			UserDB: r.PostFormValue("dbuser"),
			PwdDB: r.PostFormValue("dbpwd"),
			NameDB: r.PostFormValue("dbname"),
			HostDB: r.PostFormValue("dbhost"),
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(install, header, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}

		} else {
			echo := "echo 'DB_TEST=localhost' >> /home/tech/.profile"
			cmd := exec.Command(echo)
			out, err := cmd.Output()
			if err != nil {
				fmt.Println("Could not run command", err )
			}
			fmt.Println("Output: ", string(out))

			// data := "DB_USER="+msg.UserDB+"\nDB_PWD="+msg.PwdDB+"\nDB_NAME="+msg.NameDB+"\nDB_HOST="+msg.HostDB+"\nDB_PORT=5432"
			// env := []byte(data)
			// // env := []byte("DB_USER=tech\nDB_PWD=123456\nDB_NAME=people\nDB_HOST=localhost\n")
			// err := os.WriteFile("lib/.env", env, 0644)
			// utils.CheckError(err)

			message := "Sistema instalado"
			tmpl, _ := template.ParseFiles(ms, header, footer)
			linkmsg := "/register"
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