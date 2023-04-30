package main

import (
	"fmt"
	"net/http"
)

// maneja la ruta principal
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo!")
}

// maneja la ruta api
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo API!")
}
