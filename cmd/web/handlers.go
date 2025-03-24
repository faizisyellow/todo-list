package main

import "net/http"

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home OK"))
}
func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", nil)
}
func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SIGNUP OK"))
}
func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", nil)
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("LOGIN OK"))
}
