package main

import (
	"errors"
	"net/http"
)

func (app *AppConfig) ping(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("The server is working perfectly!\n"))
	app.serverError(w, r, errors.New("just test"));
}