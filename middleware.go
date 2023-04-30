package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// checkAuth: funcion que no tiene parametros y devuelve una funcion de tipo Middleware
func CheckAuth() Middleware {

	// Aqui estamos haciendo return nuestro Middleware
	// (Middleware es una funcion que toma http.HandlerFunc como parametro
	// y devuelve otra http.HandlerFunc)
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Como Middleware devuelve un http.HandlerFunc, aqui estamos implementando
		// ese HandlerFunc (http.HandlerFunc es una funcion que tiene
		// http.ResponseWriter y *http.Request como parametros)
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")

			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

// Logging: funcion que loguea el acceso al servidor y devuelve una funcion de tipo Middleware
func Logging() Middleware {

	// Aqui estamos haciendo return nuestro Middleware
	// (Middleware es una funcion que toma http.HandlerFunc como parametro
	// y devuelve otra http.HandlerFunc)
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Como Middleware devuelve un http.HandlerFunc, aqui estamos implementando
		// ese HandlerFunc (http.HandlerFunc es una funcion que tiene
		// http.ResponseWriter y *http.Request como parametros)
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
