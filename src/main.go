package main

import (
	"github.com/devnandito/webserver/api"
	"github.com/devnandito/webserver/handlers"
	"github.com/devnandito/webserver/middleware"
	"github.com/devnandito/webserver/server"
)

func main() {
	http := server.NewServer(":8080")
	// API clients
	http.Handle("GET", "/api/clients", api.HandleApiClients)
	http.Handle("POST", "/api/clients", api.HandleApiCreateClient)
	http.Handle("GET", "/api/clients/:id", api.HandleApiPutClient)
	http.Handle("POST", "/api/users", handlers.HandleUserPostRequest)
	http.Handle("POST", "/api/v1/users", http.AddMiddleware(handlers.HandlePostRequest, middleware.CheckAuth(), middleware.Logging()))
	
	// API module
	http.Handle("GET", "/api/modules", api.HandleApiModules)
	http.Handle("POST", "/api/modules", api.HandleApiCreateModule)

	// API operation
	http.Handle("GET", "/api/operations", api.HandleApiOperations)
	http.Handle("POST", "/api/operations", api.HandleApiCreateOperation)

	// API role
	http.Handle("GET", "/api/roles", api.HandleApiRole)
	http.Handle("POST", "/api/roles", api.HandleApiCreateRole)

	// API user
	http.Handle("GET", "/api/users", api.HandleApiUser)
	http.Handle("POST", "/api/users", api.HandleApiCreateUser)


	// TEMPLATE
	http.File("assets")
	// Register
	http.Handle("GET", "/register", handlers.SignUpUser)
	http.Handle("POST", "/register", handlers.SignUpUser)

	// Users
	http.Handle("GET", "/", handlers.SignInUser)
	http.Handle("POST", "/login", handlers.SignInUser)
	http.Handle("GET", "/logout", handlers.Logout)
	http.Handle("GET", "/users/show", http.AddMiddleware(handlers.HandelShowUser, middleware.CheckAuth()))

	// Clients
	http.Handle("GET", "/clients/show", http.AddMiddleware(handlers.HandleShowClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/create", http.AddMiddleware(handlers.HandleCreateClient, middleware.CheckAuth()))
	http.Handle("POST", "/clients/create", http.AddMiddleware(handlers.HandleCreateClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/edit", http.AddMiddleware(handlers.HandleUpdateClient, middleware.CheckAuth()))
	http.Handle("POST", "/clients/edit", http.AddMiddleware(handlers.HandleUpdateClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/detail", http.AddMiddleware(handlers.HandleGetClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/delete", http.AddMiddleware(handlers.HandleDeleteClient, middleware.CheckAuth()))
	http.Listen()
}