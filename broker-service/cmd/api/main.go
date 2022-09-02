package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", port)

	// * define http server
	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	// * start server
	err := serve.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
