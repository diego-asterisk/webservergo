package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func ConstructorServer(port string) *Server {
	return &Server{
		port:   port,
		router: ConstructorRouter(),
	}
}

func (this *Server) Listen() error {
	// http.ListenAndServe recibe el puerto y el manejador
	// pero usaremos nuestro manejador (le mandamos nil)
	http.Handle("/", this.router)
	err := http.ListenAndServe(this.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (this *Server) Handle(method string, path string, handler http.HandlerFunc) {
	this.router.addRule(method, path, handler)
}

// Devuelve el manejador resultado de aplicar los middlewares
func (this *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
