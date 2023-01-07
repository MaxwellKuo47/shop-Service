package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/maxwellkuo47/shop-Service/internal/models"
)

type AppConfig struct {
	store           *models.Store
	serverErrLogger *log.Logger
	infoLogger      *log.Logger
	clientErrLogger *log.Logger
}

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/shop_dev?sslmode=disable"
)

func main() {
	var app AppConfig
	app.serverErrLogger = log.New(os.Stdout, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
	app.infoLogger = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime)
	app.clientErrLogger = log.New(os.Stdout, "ClientERROR : ", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := initDB()
	if err != nil {
		app.serverErrLogger.Fatal(err)
	}
	app.store = models.NewStore(db)

	svr := &http.Server{
		Addr:     ":8081",
		Handler:  app.routes(),
		ErrorLog: app.serverErrLogger,
	}

	app.infoLogger.Println("Server on port :8081")
	log.Fatal(svr.ListenAndServe())
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(24)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db, err
}
