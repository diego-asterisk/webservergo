package main
// dividimos funcionalidad en server.go y otros con el mismo package para tener acceso

func main() {
	server := ConstructorServer(":3000")
	server.Handle("/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("/api", server.AddMiddleware(HandleAPI, CheckAuth(), Logging()))
	server.Listen()
}
