package main

import (
	"cv-form/config"
	"cv-form/routers"
	"net/http"
)

func main() {
	config.Connect()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routers.InitRoutes(),
	}

	server.ListenAndServe()
}
