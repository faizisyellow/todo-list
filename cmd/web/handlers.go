package main

import "net/http"

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME OK"))
}
func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SIGNUP FORM OK"))
}
func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SIGNUP OK"))
}
func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("LOGINFORM OK"))
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("LOGIN OK"))
}
