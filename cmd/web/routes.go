package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *AppConfig) routes() http.Handler {
	httpRouter := httprouter.New()
	// httpRouter.Handler(http.MethodGet, "/page",http.StripPrefix("/page/",http.FileServer(http.Dir("./build"))))
	httpRouter.HandlerFunc(http.MethodGet, "/api/ping", app.ping)
	httpRouter.HandlerFunc(http.MethodPost, "/api/product", app.productAdd)

	standard := alice.New(app.requestLogger, app.setHeader)
	return standard.Then(httpRouter)
}
