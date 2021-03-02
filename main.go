package main

func main() {
	initDB()
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/template", HandleTemplate)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}