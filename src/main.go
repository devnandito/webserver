package main

import (
	"github.com/devnandito/webserver/api"
	"github.com/devnandito/webserver/handlers"
	"github.com/devnandito/webserver/middleware"
	"github.com/devnandito/webserver/server"
)

func main() {
	http := server.NewServer(":9000")
	// API
	http.Handle("GET", "/api/clients", api.HandleApiClients)
	http.Handle("POST", "/api/clients", api.HandleApiCreateClient)
	http.Handle("GET", "/api/clients/:id", api.HandleApiPutClient)
	http.Handle("POST", "/api/users", handlers.HandleUserPostRequest)
	http.Handle("POST", "/api/v1/users", http.AddMiddleware(handlers.HandlePostRequest, middleware.CheckAuth(), middleware.Logging()))

	// TEMPLATE
	http.File("assets")
	http.Handle("GET", "/", handlers.HandleHome)
	http.Handle("GET", "/clients/show", handlers.HandleShowClient)
	// http.Handle("GET", "/clients/show", http.AddMiddleware(handlers.HandleShowClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/create", handlers.HandleCreateClient)
	http.Handle("POST", "/clients/create", handlers.HandleCreateClient)
	http.Handle("GET", "/clients/edit", handlers.HandleUpdateClient)
	http.Handle("POST", "/clients/edit", handlers.HandleUpdateClient)
	http.Handle("GET", "/clients/detail", handlers.HandleGetClient)
	http.Handle("GET", "/clients/delete", handlers.HandleDeleteClient)
	http.Listen()
	// http.Handle("GET", "/root", handlers.HandleRoot)
}