package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}

	db.Close()

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to todo-list web app."))
	}))

	infoLog.Printf("Starting server on %s", ADDRESS)
	err = http.ListenAndServe(ADDRESS, mux)
	errorLog.Fatal(err)
}

func openDB() (*sql.DB, error) {
	dns := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
