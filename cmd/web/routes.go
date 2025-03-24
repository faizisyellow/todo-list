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
	router.HandleFunc("/login", app.loginForm).Methods("GET")
	router.HandleFunc("/login", app.login).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	router.HandleFunc("/ping", app.ping)
	return router
}
