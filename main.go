package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	r := RegisterRoutes()

	// use handlers as logger
	r.Use(AccessLogMiddleware)
	r.Use(ContentNegotiatorMiddleware)
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(false)))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
