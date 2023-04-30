package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string] map[string]http.HandlerFunc
}

func ConstructorRouter() *Router {
	return &Router{
		rules: make(map[string] map[string]http.HandlerFunc),
	}
}

//completar la interface de Router
func (this *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists, methodExists := this.FindHandler(request.Method, request.URL.Path)
	if !exists {
		fmt.Println("Path no existe")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExists {
		fmt.Println("Path existe pero Metodo no existe")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Path existe")
	handler(w, request)
}

// devuelve la llave del map, y un boll que dice si existe o no existe
func (this *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	// buscar en un map una clave y saber si existe
	_, exists := this.rules[path]
	handler, methodExists := this.rules[path][method]
	return handler, exists, methodExists
}

func (this *Router) addRule(method string, path string, handler http.HandlerFunc){
	_, exists := this.rules[path]
	if !exists {
		this.rules[path] = make(map[string]http.HandlerFunc)
	}
	this.rules[path][method] = handler
}