package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("broker service listening on port %s\n", listenPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", listenPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	log.Println(err)
}
