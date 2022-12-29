package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *AppConfig) routes() *httprouter.Router {
	httpRouter := httprouter.New()
	// httpRouter.Handler(http.MethodGet, "/page",http.StripPrefix("/page/",http.FileServer(http.Dir("./build"))))
	httpRouter.HandlerFunc(http.MethodGet, "/api/ping", app.requestLogger(app.ping))
	httpRouter.HandlerFunc(http.MethodPost, "/api/product", app.requestLogger(app.productAdd))

	return httpRouter
}
