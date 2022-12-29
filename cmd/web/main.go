package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

type AppConfig struct {
	DB              *sql.DB
	serverErrLogger *log.Logger
	infoLogger      *log.Logger
	clientErrLogger *log.Logger
}

func main() {
	var app AppConfig
	app.serverErrLogger = log.New(os.Stdout, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
	app.infoLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime)
	app.clientErrLogger = log.New(os.Stdout, "ClientERROR : ", log.Ldate|log.Ltime|log.Lshortfile)

	svr := &http.Server{
		Addr:     ":8081",
		Handler:  app.routes(),
		ErrorLog: app.serverErrLogger,
	}

	log.Fatal(svr.ListenAndServe())
}
