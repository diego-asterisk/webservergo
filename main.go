package main

import "net/http"

// dividimos funcionalidad en server.go y otros con el mismo package para tener acceso

func main() {
	server := ConstructorServer(":3000")
	server.Handle(http.MethodGet, "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle(http.MethodPost, "/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Listen()
}
