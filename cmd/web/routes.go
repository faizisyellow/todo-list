package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home).Methods("GET")
	router.HandleFunc("/signup", app.signupForm).Methods("GET")
	router.HandleFunc("/signup", app.signup).Methods("POST")
	router.HandleFunc("/login", app.login).Methods("GET")
	router.HandleFunc("/login", app.loginForm).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/ping", app.ping)
	return router
}
