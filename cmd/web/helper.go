package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *AppConfig) clientError(w http.ResponseWriter, r *http.Request, status int, logString string) {
	outputText := http.StatusText(status)
	app.clientErrLogger.Printf("ClientError => %q -> %q response '%d %s' : %q", r.Method, r.URL.Path, status, outputText, logString)
	http.Error(w, outputText, status)
}

func (app *AppConfig) serverError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("ServerError => %q -> %q : %q\n%s", r.Method, r.URL.Path, err.Error(), debug.Stack())
	app.serverErrLogger.Output(2, trace)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}