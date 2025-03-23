package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var ADDRESS = "localhost:8000"

func main() {
	infoLog := log.New(os.Stdout, "\033[32mINFO\t\033[0m", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "\033[31mERROR\t\033[0m", log.Ldate|log.Ltime|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		errorLog.Printf("can't load env file: %q", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to todo-list web app."))
	}))

	infoLog.Printf("Starting server on %s", ADDRESS)
	err = http.ListenAndServe(ADDRESS, mux)
	errorLog.Fatal(err)
}
