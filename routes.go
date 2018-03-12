package main

import (
	"go-boilerplate/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// add route
	r.HandleFunc("/", handler.TodoHandler)
	http.Handle("/", r)

	return r
}
