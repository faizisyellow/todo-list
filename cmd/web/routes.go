package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.logRequest, app.secureHeaders, app.recoverPanic)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	router := mux.NewRouter()

	router.Handle("/", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.home)).Methods("GET")
	router.Handle("/todos", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.newTodo)).Methods("POST")
	router.Handle("/todos/{id}", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.detailTodo)).Methods("GET")
	router.Handle("/todos/{actions}/{id}", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.actionsTodo)).Methods("POST")

	router.Handle("/signup", dynamicMiddleware.Append(app.userAlreadyAuthenticated).ThenFunc(app.signupForm)).Methods("GET")
	router.Handle("/signup", dynamicMiddleware.Append(app.userAlreadyAuthenticated).ThenFunc(app.signup)).Methods("POST")
	router.Handle("/login", dynamicMiddleware.Append(app.userAlreadyAuthenticated).ThenFunc(app.loginForm)).Methods("GET")
	router.Handle("/login", dynamicMiddleware.Append(app.userAlreadyAuthenticated).ThenFunc(app.login)).Methods("POST")

	router.Handle("/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logout)).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/"))))

	router.HandleFunc("/ping", app.ping)
	return standardMiddleware.Then(router)
}
