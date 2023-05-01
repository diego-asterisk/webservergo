package main

import (
	"encoding/json"
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

// maneja el primer Post
func HandlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetadataUser
	error := decoder.Decode(&metadata)
	if error != nil {
		fmt.Fprintf(w, "Error %+v", error)
		return
	}
	fmt.Fprintf(w, "Hola Metadata!")
	//fmt.Fprintln(metadata["name"])
	fmt.Fprintf(w, "Payload %#v", metadata)
}

// maneja el User Post
func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata User
	error := decoder.Decode(&metadata)
	if error != nil {
		fmt.Fprintf(w, "Error %+v", error)
		return
	}
	response, err := metadata.toJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
