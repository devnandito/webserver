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

var cont models.Contribution

func HandleShowContribution(w http.ResponseWriter, r *http.Request) {
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "List contribution"
	headers := [7]string{"ID", "Amount", "Contribution date", "Method pay", "Product", "Description", "Client"}
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	show := filepath.Join("views/contributions", "show.html")
	response, err := cont.ShowContributionGorm()
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

func HandleCreateContribution(w http.ResponseWriter, r *http.Request) {
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	m := utils.GetMenu()
	title := "Add contribution"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	add := filepath.Join("views/contributions", "add.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[7]

	clients, err := cls.ShowClientGorm()
	if err != nil {
		log.Println(err)
	}

	products, err := pro.ShowProductGorm()
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
			"Clients": clients,
			"Products": products,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template:", res)
			return
		}
	case "POST":
		contDate := utils.StrToTime(r.PostFormValue("contribution_date"))
		intAmount := utils.StrToInt(r.PostFormValue("amount"))
		intClient := utils.StrToInt(r.PostFormValue("clientid"))
		intProduct := utils.StrToInt(r.PostFormValue("productid"))
		msg := &utils.ValidateContribution{
			Amount: intAmount,
			Contribution_date: contDate,
			Method_pay: r.PostFormValue("method_pay"),
			Description: r.PostFormValue("description"),
			ClientID: intClient,
			ProductID: intProduct,
		}
		
		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(add, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"UserSession": userSession,
				"Clients": clients,
				"Products": products,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Contribution {
				Amount: msg.Amount,
				Contribution_date: msg.Contribution_date,
				Method_pay: msg.Method_pay,
				Description: msg.Description,
				ClientID: msg.ClientID,
				ProductID: msg.ProductID,
			}
			
			response, err := cont.CreateContributionGorm(&data)
	
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

func HandleUpdateContribution(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Edit contribution"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	edit := filepath.Join("views/contributions", "edit.html")
	ms := filepath.Join("views/messages", "message.html")
	url := m[7]
	products, err := pro.ShowProductGorm()
	if err != nil {
		log.Println(err)
	}

	clients, err := cls.ShowClientGorm()
	if err != nil {
		log.Println(err)
	}
	
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
		sid := r.URL.Query().Get("id")
		id, err :=  strconv.Atoi(sid)

		if err != nil {
			panic(err)
		}

		response, err := cont.GetOneContributionGorm(id)
		if err != nil {
			log.Println("Error executing template", response)
		}

		msg := &utils.ValidateContribution {
			Amount: response.Amount,
			Contribution_date: response.Contribution_date,
			Method_pay: response.Method_pay,
			Description: response.Description,
			ClientID: response.ClientID,
			ProductID: response.ProductID,
		}

		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Msg": msg,
			"ID": id,
			"UserSession": userSession,
			"RoleSession": roleSession,
			"Products": products,
			"Clients": clients,
			"FKC": response.ClientID,
			"FKP": response.ProductID,
			"Menu": m,
		})
	
		if res != nil {
			log.Println("Error executing template", res)
			return
		}

	case "POST":
		intAmount := utils.StrToInt(r.PostFormValue("amount"))
		ttime := utils.StrToTime(r.PostFormValue("contribution_date"))
		intClient := utils.StrToInt(r.PostFormValue("clientid"))
		intProduct := utils.StrToInt(r.PostFormValue("productid"))
		msg := &utils.ValidateContribution {
			Amount: intAmount,
			Contribution_date: ttime,
			Method_pay: r.PostFormValue("method_pay"),
			Description: r.PostFormValue("description"),
			ClientID: intClient,
			ProductID: intProduct,
		}

		sid := r.PostFormValue("id")
		id, err := strconv.Atoi(sid)
		
		if err != nil {
			panic(err)
		}

		clsid, err := strconv.Atoi(r.PostFormValue("clientid"))
		if err != nil {
			log.Println(err)
		}

		pid, err := strconv.Atoi(r.PostFormValue("productid"))
		if err != nil {
			log.Println(err)
		}

		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(edit, header, nav, menu, javascript, footer)
			res := tmpl.Execute(w, map[string]interface{}{
				"Title": title,
				"Msg": msg,
				"ID": id,
				"UserSession": userSession,
				"RoleSession": roleSession,
				"Products": products,
				"Clients": clients,
				"FKC": clsid,
				"FKP": pid,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template", res)
				return
			}
		} else {
			data := models.Contribution {
				Amount: msg.Amount,
				Contribution_date: msg.Contribution_date,
				Method_pay: msg.Method_pay,
				Description: msg.Description,
				ClientID: msg.ClientID,
				ProductID: msg.ProductID,
			}
			
			response, err := cont.UpdateContributionGorm(id, &data)
	
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

func HandleGetContribution(w http.ResponseWriter, r *http.Request){
	m := utils.GetMenu()
	session := utils.GetSession(r)
	userSession := session.Values["username"]
	roleSession := session.Values["role"]
	title := "Delete contribution"
	header := filepath.Join("views", "header.html")
	nav := filepath.Join("views", "nav.html")
	menu := filepath.Join("views", "menu.html")
	javascript := filepath.Join("views", "javascript.html")
	footer := filepath.Join("views", "footer.html")
	delete := filepath.Join("views/contributions", "delete.html")
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		panic(err)
	}
	response, err := cont.GetOneContributionGorm(id)
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

func HandleDeleteContribution(w http.ResponseWriter, r *http.Request){
	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)

	if err != nil {
		panic(err)
	}

	response := cont.DeleteContributionGorm(id)
	log.Println("Deleted contribution", response)
	http.Redirect(w, r, "/contributions/show", http.StatusFound)
}