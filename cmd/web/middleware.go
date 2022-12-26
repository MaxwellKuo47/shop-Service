package main

import "net/http"

func (app *AppConfig)requestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.infoLogger.Printf(" %s -> %q", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}