package main

import (
	"github.com/devnandito/webserver/api"
	"github.com/devnandito/webserver/handlers"
	"github.com/devnandito/webserver/middleware"
	"github.com/devnandito/webserver/server"
)

func main() {
	server := server.NewServer(":9000")
	// API
	server.Handle("POST", "/api/clients", server.AddMiddleware(handlers.HandlePostRequest, middleware.CheckAuth(), middleware.Logging()))
	server.Handle("GET", "/api/clients", api.HandleApiClientGet)
	server.Handle("POST", "/api/users", handlers.HandleUserPostRequest)
	// TEMPLATE
	server.Handle("GET", "/", handlers.HandleHome)
	server.Handle("GET", "/clients/show", handlers.HandleShowClient)
	server.Handle("GET", "/root", handlers.HandleRoot)
	server.Listen()
}