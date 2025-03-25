package main

import (
	"net/http"

	"faizisyellow.com/todolist/pkg/forms"
	"faizisyellow.com/todolist/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) newTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	user := app.authenticatedUser(r)

	form := forms.New(r.PostForm)
	form.Required("task")
	form.MinLength("task", 3)

	if !form.Valid() {
		app.render(w, r, "home.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	err = app.todos.Insert(form.Get("task"), user.ID)
	if err == models.ErrRequireUser {
		form.Errors.Add("task", "Invalid input")
		app.render(w, r, "home.page.tmpl", &templateData{
			Form: form,
		})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) signupForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	err = app.users.Insert(form.Get("email"), form.Get("name"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "address already in use")
		app.render(w, r, "signup.page.tmpl", &templateData{
			Form: form,
		})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) loginForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("email", "password")

	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Email or Password is incorrect")
		app.render(w, r, "login.page.tmpl", &templateData{
			Form: form,
		})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "userID", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "userID")

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
