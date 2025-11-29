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


	// Install
	http.Handle("GET", "/install", handlers.HandleInstall)
	http.Handle("POST", "/install", handlers.HandleInstall)
	http.Handle("GET", "/seed", handlers.HandleSeedData)
	http.Handle("POST", "/seed", handlers.HandleSeedData)
	
	// Users
	http.Handle("GET", "/", handlers.SignInUser)
	http.Handle("POST", "/login", handlers.SignInUser)
	http.Handle("GET", "/logout", handlers.Logout)
	http.Handle("GET", "/users/show", http.AddMiddleware(handlers.HandleShowUser, middleware.CheckAuth()))
	http.Handle("GET", "/users/create", http.AddMiddleware(handlers.HandleCreateUser, middleware.CheckAuth()))
	http.Handle("POST", "/users/create", http.AddMiddleware(handlers.HandleCreateUser, middleware.CheckAuth()))
	http.Handle("GET", "/users/put", http.AddMiddleware(handlers.HandleUpdateUser, middleware.CheckAuth()))
	http.Handle("POST", "/users/put", http.AddMiddleware(handlers.HandleUpdateUser, middleware.CheckAuth()))
	http.Handle("GET", "/users/detail", http.AddMiddleware(handlers.HandleGetUser, middleware.CheckAuth()))
	http.Handle("GET", "/users/delete", http.AddMiddleware(handlers.HandleDeleteUser, middleware.CheckAuth()))
	http.Handle("GET", "/users/change", http.AddMiddleware(handlers.HandleChangePwd, middleware.CheckAuth()))
	http.Handle("POST", "/users/change", http.AddMiddleware(handlers.HandleChangePwd, middleware.CheckAuth()))
	http.Handle("GET", "/dashboard", http.AddMiddleware(handlers.HandleDashboard, middleware.CheckAuth()))
	http.Handle("GET", "/profiles/show", http.AddMiddleware(handlers.HandleProfileUser, middleware.CheckAuth()))
	http.Handle("GET", "/profiles/put", http.AddMiddleware(handlers.HandleUpdateProfile, middleware.CheckAuth()))
	http.Handle("POST", "/profiles/put", http.AddMiddleware(handlers.HandleUpdateProfile, middleware.CheckAuth()))
	http.Handle("GET", "/profiles/change", http.AddMiddleware(handlers.HandleChangeProfilePwd, middleware.CheckAuth()))
	http.Handle("POST", "/profiles/change", http.AddMiddleware(handlers.HandleChangeProfilePwd, middleware.CheckAuth()))
	
	// Modules
	http.Handle("GET", "/modules/show", http.AddMiddleware(handlers.HandelShowModule, middleware.CheckAuth()))
	http.Handle("GET", "/modules/create", http.AddMiddleware(handlers.HandleCreateModule, middleware.CheckAuth()))
	http.Handle("POST", "/modules/create", http.AddMiddleware(handlers.HandleCreateModule, middleware.CheckAuth()))
	http.Handle("GET", "/modules/put", http.AddMiddleware(handlers.HandleUpdateModule, middleware.CheckAuth()))
	http.Handle("POST", "/modules/put", http.AddMiddleware(handlers.HandleUpdateModule, middleware.CheckAuth()))
	http.Handle("GET", "/modules/detail", http.AddMiddleware(handlers.HandleGetModule, middleware.CheckAuth()))
	http.Handle("GET", "/modules/delete", http.AddMiddleware(handlers.HandleDeleteModule, middleware.CheckAuth()))
	
	// Operations
	http.Handle("GET", "/operations/show", http.AddMiddleware(handlers.HandelShowOperation, middleware.CheckAuth()))
	http.Handle("GET", "/operations/create", http.AddMiddleware(handlers.HandleCreateOperation, middleware.CheckAuth()))
	http.Handle("POST", "/operations/create", http.AddMiddleware(handlers.HandleCreateOperation, middleware.CheckAuth()))
	http.Handle("GET", "/operations/put", http.AddMiddleware(handlers.HandleUpdateOperation, middleware.CheckAuth()))
	http.Handle("POST", "/operations/put", http.AddMiddleware(handlers.HandleUpdateOperation, middleware.CheckAuth()))
	http.Handle("GET", "/operations/detail", http.AddMiddleware(handlers.HandleGetOperation, middleware.CheckAuth()))
	http.Handle("GET", "/operations/delete", http.AddMiddleware(handlers.HandleDeleteOperation, middleware.CheckAuth()))
	
	// Roles
	http.Handle("GET", "/roles/show", http.AddMiddleware(handlers.HandelShowRole, middleware.CheckAuth()))
	http.Handle("GET", "/roles/create", http.AddMiddleware(handlers.HandleCreateRole, middleware.CheckAuth()))
	http.Handle("POST", "/roles/create", http.AddMiddleware(handlers.HandleCreateRole, middleware.CheckAuth()))
	http.Handle("GET", "/roles/put", http.AddMiddleware(handlers.HandleUpdateRole, middleware.CheckAuth()))
	http.Handle("POST", "/roles/put", http.AddMiddleware(handlers.HandleUpdateRole, middleware.CheckAuth()))
	http.Handle("GET", "/roles/detail", http.AddMiddleware(handlers.HandleGetRole, middleware.CheckAuth()))
	http.Handle("GET", "/roles/delete", http.AddMiddleware(handlers.HandleDeleteRole, middleware.CheckAuth()))

	// Clients
	http.Handle("GET", "/clients/show", http.AddMiddleware(handlers.HandleShowClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/create", http.AddMiddleware(handlers.HandleCreateClient, middleware.CheckAuth()))
	http.Handle("POST", "/clients/create", http.AddMiddleware(handlers.HandleCreateClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/put", http.AddMiddleware(handlers.HandleUpdateClient, middleware.CheckAuth()))
	http.Handle("POST", "/clients/put", http.AddMiddleware(handlers.HandleUpdateClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/detail", http.AddMiddleware(handlers.HandleGetClient, middleware.CheckAuth()))
	http.Handle("GET", "/clients/delete", http.AddMiddleware(handlers.HandleDeleteClient, middleware.CheckAuth()))

	// Products
	http.Handle("GET", "/products/show", http.AddMiddleware(handlers.HandleShowProduct, middleware.CheckAuth()))
	http.Handle("GET", "/products/create", http.AddMiddleware(handlers.HandleCreateProduct, middleware.CheckAuth()))
	http.Handle("POST", "/products/create", http.AddMiddleware(handlers.HandleCreateProduct, middleware.CheckAuth()))
	http.Handle("GET", "/products/put", http.AddMiddleware(handlers.HandleUpdateProduct, middleware.CheckAuth()))
	http.Handle("POST", "/products/put", http.AddMiddleware(handlers.HandleUpdateProduct, middleware.CheckAuth()))
	http.Handle("GET", "/products/detail", http.AddMiddleware(handlers.HandleGetProduct, middleware.CheckAuth()))
	http.Handle("GET", "/products/delete", http.AddMiddleware(handlers.HandleDeleteProduct, middleware.CheckAuth()))

	// Contributions
	http.Handle("GET", "/contributions/show", http.AddMiddleware(handlers.HandleShowContribution, middleware.CheckAuth()))
	http.Handle("GET", "/contributions/create", http.AddMiddleware(handlers.HandleCreateContribution, middleware.CheckAuth()))
	http.Handle("POST", "/contributions/create", http.AddMiddleware(handlers.HandleCreateContribution, middleware.CheckAuth()))
	http.Handle("GET", "/contributions/put", http.AddMiddleware(handlers.HandleUpdateContribution, middleware.CheckAuth()))
	http.Handle("POST", "/contributions/put", http.AddMiddleware(handlers.HandleUpdateContribution, middleware.CheckAuth()))
	http.Handle("GET", "/contributions/detail", http.AddMiddleware(handlers.HandleGetContribution, middleware.CheckAuth()))
	http.Handle("GET", "/contributions/delete", http.AddMiddleware(handlers.HandleDeleteContribution, middleware.CheckAuth()))

	http.Listen()
}