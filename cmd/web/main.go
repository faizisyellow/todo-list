package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"faizisyellow.com/todolist/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var ADDRESS = "localhost:4000"

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	users    *mysql.UserModel
}

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

	// app := application{
	// 	infoLog:  infoLog,
	// 	errorLog: errorLog,
	// 	users: &mysql.UserModel{
	// 		DB: db,
	// 	},
	// }

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         ADDRESS,
		ErrorLog:     errorLog,
		Handler:      mux,
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", ADDRESS)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
