package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devnandito/webserver/models"
	"github.com/devnandito/webserver/utils"
)

func HandleSeedData(w http.ResponseWriter, r *http.Request){
	title := "Seed"
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	install := filepath.Join("views/install", "seed.html")
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
		msg := &utils.ValidateSeed{
			ModuleDes: r.PostFormValue("module"),
			OperationDes: r.PostFormValue("operation"),
			RoleDes: r.PostFormValue("role"),
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
			dataModule := models.Module {
				Description: msg.ModuleDes,
			}
			module, err := mod.CreateModuleGorm(&dataModule)
			utils.CheckError(err)
			fmt.Println(module)

			dataOperation := models.Operation {
				Description: msg.OperationDes,
				ModuleID: 1,
			}
			operation, err := op.CreateOperationGorm(&dataOperation)
			utils.CheckError(err)
			fmt.Println(operation)

			dataRole := models.Role {
				Description: msg.RoleDes,
			}
			role, err := rol.CreateRoleGorm(&dataRole)
			utils.CheckError(err)
			fmt.Println(role)

			message := "Datos insertados"
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