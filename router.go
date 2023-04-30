package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func ConstructorRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

//completar la interface de Router
func (this *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists := this.FindHandler(request.URL.Path)
	if !exists {
		fmt.Println("Path no existe")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Path existe")
	handler(w, request)
}

// devuelve la llave del map, y un boll que dice si existe o no existe
func (this *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	// buscar en un map una clave y saber si existe
	handler, exists := this.rules[path]
	return handler, exists
}

func (this *Router) addRule(path string, handler http.HandlerFunc){
	this.rules[path] = handler
}