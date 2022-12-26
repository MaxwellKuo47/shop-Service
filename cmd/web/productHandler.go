package main

import "net/http"

const MAX_FORM_ACCEPT = 10 * 1024 * 1024

func (app *AppConfig) productAdd (w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(MAX_FORM_ACCEPT) //10MB
} 